package cloudwatch

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grafana/grafana/pkg/tsdb/cloudwatch/mocks"
	"github.com/grafana/grafana/pkg/tsdb/cloudwatch/models"
	"github.com/grafana/grafana/pkg/tsdb/cloudwatch/models/resources"
	"github.com/grafana/grafana/pkg/tsdb/cloudwatch/services"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Metrics_Route(t *testing.T) {
	origNewListMetricsServices := services.NewListMetricsService
	t.Cleanup(func() {
		services.NewListMetricsService = origNewListMetricsServices
	})
	var mockListMetricsService mocks.ListMetricsServiceMock
	services.NewListMetricsService = func(provider models.MetricsClientProvider) models.ListMetricsProvider {
		return &mockListMetricsService
	}
	t.Run("calls GetMetricsByNamespace when a CustomNamespaceRequestType is passed", func(t *testing.T) {
		mockListMetricsService = mocks.ListMetricsServiceMock{}
		mockListMetricsService.On("GetMetricsByNamespace", mock.Anything).Return([]resources.ResourceResponse[resources.Metric]{}, nil)
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/metrics?region=us-east-2&namespace=customNamespace", nil)
		ds := newTestDatasource()
		handler := http.HandlerFunc(ds.resourceRequestMiddleware(ds.MetricsHandler))
		handler.ServeHTTP(rr, req)
		mockListMetricsService.AssertNumberOfCalls(t, "GetMetricsByNamespace", 1)
	})

	t.Run("calls GetAllHardCodedMetrics when a AllMetricsRequestType is passed", func(t *testing.T) {
		origGetAllHardCodedMetrics := services.GetAllHardCodedMetrics
		t.Cleanup(func() {
			services.GetAllHardCodedMetrics = origGetAllHardCodedMetrics
		})
		haveBeenCalled := false
		services.GetAllHardCodedMetrics = func() []resources.ResourceResponse[resources.Metric] {
			haveBeenCalled = true
			return []resources.ResourceResponse[resources.Metric]{}
		}
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/metrics?region=us-east-2", nil)
		ds := newTestDatasource()
		handler := http.HandlerFunc(ds.resourceRequestMiddleware(ds.MetricsHandler))
		handler.ServeHTTP(rr, req)
		assert.True(t, haveBeenCalled)
	})

	t.Run("calls GetHardCodedMetricsByNamespace when a MetricsByNamespaceRequestType is passed", func(t *testing.T) {
		origGetHardCodedMetricsByNamespace := services.GetHardCodedMetricsByNamespace
		t.Cleanup(func() {
			services.GetHardCodedMetricsByNamespace = origGetHardCodedMetricsByNamespace
		})
		haveBeenCalled := false
		usedNamespace := ""
		services.GetHardCodedMetricsByNamespace = func(namespace string) ([]resources.ResourceResponse[resources.Metric], error) {
			haveBeenCalled = true
			usedNamespace = namespace
			return []resources.ResourceResponse[resources.Metric]{}, nil
		}
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/metrics?region=us-east-2&namespace=AWS/DMS", nil)
		ds := newTestDatasource()
		handler := http.HandlerFunc(ds.resourceRequestMiddleware(ds.MetricsHandler))
		handler.ServeHTTP(rr, req)
		assert.True(t, haveBeenCalled)
		assert.Equal(t, "AWS/DMS", usedNamespace)
	})

	t.Run("returns 500 if GetMetricsByNamespace returns an error", func(t *testing.T) {
		mockListMetricsService = mocks.ListMetricsServiceMock{}
		mockListMetricsService.On("GetMetricsByNamespace", mock.Anything).Return([]resources.ResourceResponse[resources.Metric]{}, fmt.Errorf("some error"))
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/metrics?region=us-east-2&namespace=customNamespace", nil)
		ds := newTestDatasource()
		handler := http.HandlerFunc(ds.resourceRequestMiddleware(ds.MetricsHandler))
		handler.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusInternalServerError, rr.Code)
		assert.Equal(t, `{"Message":"error in MetricsHandler: some error","Error":"some error","StatusCode":500}`, rr.Body.String())
	})
}
