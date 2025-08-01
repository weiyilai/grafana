package featuremgmt

import (
	"context"
	"fmt"

	"github.com/grafana/grafana/pkg/infra/log"
	"github.com/grafana/grafana/pkg/setting"
	"github.com/open-feature/go-sdk/openfeature"
	goffmodel "github.com/thomaspoignant/go-feature-flag/cmd/relayproxy/model"
)

// StaticFlagEvaluator provides methods for evaluating static feature flags
// it is only used when static provider is configured
type StaticFlagEvaluator interface {
	EvalFlag(ctx context.Context, flagKey string) (goffmodel.OFREPEvaluateSuccessResponse, error)
	EvalAllFlags(ctx context.Context) (goffmodel.OFREPBulkEvaluateSuccessResponse, error)
}

// CreateStaticEvaluator is a dependency for ofrep APIBuilder
func CreateStaticEvaluator(cfg *setting.Cfg) (StaticFlagEvaluator, error) {
	if cfg.OpenFeature.ProviderType != setting.StaticProviderType {
		return nil, fmt.Errorf("provider is not a static provider, type %s", setting.StaticProviderType)
	}

	staticFlags, err := setting.ReadFeatureTogglesFromInitFile(cfg.Raw.Section("feature_toggles"))
	if err != nil {
		return nil, fmt.Errorf("failed to read feature flags from config: %w", err)
	}

	staticProvider, err := newStaticProvider(staticFlags)
	if err != nil {
		return nil, fmt.Errorf("failed to create static provider: %w", err)
	}

	p, ok := staticProvider.(*inMemoryBulkProvider)
	if !ok {
		return nil, fmt.Errorf("static provider is not of type inMemoryBulkProvider")
	}

	c := openfeature.GetApiInstance().GetClient()

	return &staticEvaluator{
		provider: p,
		client:   c,
		log:      log.New("static-evaluator"),
	}, nil
}

// staticEvaluator implements StaticFlagEvaluator for static providers
type staticEvaluator struct {
	provider *inMemoryBulkProvider
	client   openfeature.IClient
	log      log.Logger
}

func (s *staticEvaluator) EvalFlag(ctx context.Context, flagKey string) (goffmodel.OFREPEvaluateSuccessResponse, error) {
	result, err := s.client.BooleanValueDetails(ctx, flagKey, false, openfeature.TransactionContext(ctx))
	if err != nil {
		return goffmodel.OFREPEvaluateSuccessResponse{}, fmt.Errorf("failed to evaluate flag %s: %w", flagKey, err)
	}
	resp := goffmodel.OFREPEvaluateSuccessResponse{
		Key:      flagKey,
		Value:    result.Value,
		Reason:   "static provider evaluation result",
		Variant:  result.Variant,
		Metadata: result.FlagMetadata,
	}

	return resp, nil
}

func (s *staticEvaluator) EvalAllFlags(ctx context.Context) (goffmodel.OFREPBulkEvaluateSuccessResponse, error) {
	flags, err := s.provider.ListFlags()
	if err != nil {
		return goffmodel.OFREPBulkEvaluateSuccessResponse{}, fmt.Errorf("static provider failed to list all flags: %w", err)
	}

	allFlags := make([]goffmodel.OFREPFlagBulkEvaluateSuccessResponse, 0, len(flags))
	for _, flagKey := range flags {
		result, err := s.client.BooleanValueDetails(ctx, flagKey, false, openfeature.TransactionContext(ctx))
		if err != nil {
			s.log.Error("failed to evaluate flag during bulk evaluation", "flagKey", flagKey, "error", err)
			continue
		}

		allFlags = append(allFlags, goffmodel.OFREPFlagBulkEvaluateSuccessResponse{
			OFREPEvaluateSuccessResponse: goffmodel.OFREPEvaluateSuccessResponse{
				Key:      flagKey,
				Value:    result.Value,
				Reason:   "static provider evaluation result",
				Variant:  result.Variant,
				Metadata: result.FlagMetadata,
			},
			ErrorCode:    string(result.ErrorCode),
			ErrorDetails: result.ErrorMessage,
		})
	}

	return goffmodel.OFREPBulkEvaluateSuccessResponse{Flags: allFlags}, nil
}
