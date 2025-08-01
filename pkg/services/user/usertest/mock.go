// Code generated by mockery v2.53.4. DO NOT EDIT.

package usertest

import (
	context "context"

	user "github.com/grafana/grafana/pkg/services/user"
	mock "github.com/stretchr/testify/mock"
)

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

// BatchDisableUsers provides a mock function with given fields: _a0, _a1
func (_m *MockService) BatchDisableUsers(_a0 context.Context, _a1 *user.BatchDisableUsersCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for BatchDisableUsers")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.BatchDisableUsersCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *MockService) Create(_a0 context.Context, _a1 *user.CreateUserCommand) (*user.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.CreateUserCommand) (*user.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.CreateUserCommand) *user.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.CreateUserCommand) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateServiceAccount provides a mock function with given fields: _a0, _a1
func (_m *MockService) CreateServiceAccount(_a0 context.Context, _a1 *user.CreateUserCommand) (*user.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for CreateServiceAccount")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.CreateUserCommand) (*user.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.CreateUserCommand) *user.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.CreateUserCommand) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *MockService) Delete(_a0 context.Context, _a1 *user.DeleteUserCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.DeleteUserCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByEmail provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetByEmail(_a0 context.Context, _a1 *user.GetUserByEmailQuery) (*user.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetByEmail")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByEmailQuery) (*user.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByEmailQuery) *user.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetUserByEmailQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetByID(_a0 context.Context, _a1 *user.GetUserByIDQuery) (*user.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetByID")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByIDQuery) (*user.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByIDQuery) *user.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetUserByIDQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByLogin provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetByLogin(_a0 context.Context, _a1 *user.GetUserByLoginQuery) (*user.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetByLogin")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByLoginQuery) (*user.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByLoginQuery) *user.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetUserByLoginQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUID provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetByUID(_a0 context.Context, _a1 *user.GetUserByUIDQuery) (*user.User, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetByUID")
	}

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByUIDQuery) (*user.User, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserByUIDQuery) *user.User); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetUserByUIDQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProfile provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetProfile(_a0 context.Context, _a1 *user.GetUserProfileQuery) (*user.UserProfileDTO, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetProfile")
	}

	var r0 *user.UserProfileDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserProfileQuery) (*user.UserProfileDTO, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetUserProfileQuery) *user.UserProfileDTO); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.UserProfileDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetUserProfileQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSignedInUser provides a mock function with given fields: _a0, _a1
func (_m *MockService) GetSignedInUser(_a0 context.Context, _a1 *user.GetSignedInUserQuery) (*user.SignedInUser, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetSignedInUser")
	}

	var r0 *user.SignedInUser
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetSignedInUserQuery) (*user.SignedInUser, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.GetSignedInUserQuery) *user.SignedInUser); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.SignedInUser)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.GetSignedInUserQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsageStats provides a mock function with given fields: ctx
func (_m *MockService) GetUsageStats(ctx context.Context) map[string]interface{} {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetUsageStats")
	}

	var r0 map[string]interface{}
	if rf, ok := ret.Get(0).(func(context.Context) map[string]interface{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	return r0
}

// ListByIdOrUID provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockService) ListByIdOrUID(_a0 context.Context, _a1 []string, _a2 []int64) ([]*user.User, error) {
	ret := _m.Called(_a0, _a1, _a2)

	if len(ret) == 0 {
		panic("no return value specified for ListByIdOrUID")
	}

	var r0 []*user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, []int64) ([]*user.User, error)); ok {
		return rf(_a0, _a1, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string, []int64) []*user.User); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string, []int64) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Search provides a mock function with given fields: _a0, _a1
func (_m *MockService) Search(_a0 context.Context, _a1 *user.SearchUsersQuery) (*user.SearchUserQueryResult, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Search")
	}

	var r0 *user.SearchUserQueryResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.SearchUsersQuery) (*user.SearchUserQueryResult, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *user.SearchUsersQuery) *user.SearchUserQueryResult); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.SearchUserQueryResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *user.SearchUsersQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *MockService) Update(_a0 context.Context, _a1 *user.UpdateUserCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.UpdateUserCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateLastSeenAt provides a mock function with given fields: _a0, _a1
func (_m *MockService) UpdateLastSeenAt(_a0 context.Context, _a1 *user.UpdateUserLastSeenAtCommand) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateLastSeenAt")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *user.UpdateUserLastSeenAtCommand) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
