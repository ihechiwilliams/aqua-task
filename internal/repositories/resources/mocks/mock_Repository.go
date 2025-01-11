// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	resources "aqua-backend/internal/repositories/resources"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// CreateResourcesByCustomerID provides a mock function with given fields: ctx, customerID, _a2
func (_m *MockRepository) CreateResourcesByCustomerID(ctx context.Context, customerID uuid.UUID, _a2 []*resources.DBResource) ([]*resources.Resource, error) {
	ret := _m.Called(ctx, customerID, _a2)

	if len(ret) == 0 {
		panic("no return value specified for CreateResourcesByCustomerID")
	}

	var r0 []*resources.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []*resources.DBResource) ([]*resources.Resource, error)); ok {
		return rf(ctx, customerID, _a2)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, []*resources.DBResource) []*resources.Resource); ok {
		r0 = rf(ctx, customerID, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*resources.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, []*resources.DBResource) error); ok {
		r1 = rf(ctx, customerID, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteResource provides a mock function with given fields: ctx, id
func (_m *MockRepository) DeleteResource(ctx context.Context, id uuid.UUID) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteResource")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetResourceByID provides a mock function with given fields: ctx, id
func (_m *MockRepository) GetResourceByID(ctx context.Context, id uuid.UUID) (*resources.Resource, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetResourceByID")
	}

	var r0 *resources.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*resources.Resource, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *resources.Resource); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*resources.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourcesByCustomerID provides a mock function with given fields: ctx, customerID
func (_m *MockRepository) GetResourcesByCustomerID(ctx context.Context, customerID uuid.UUID) ([]*resources.Resource, error) {
	ret := _m.Called(ctx, customerID)

	if len(ret) == 0 {
		panic("no return value specified for GetResourcesByCustomerID")
	}

	var r0 []*resources.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) ([]*resources.Resource, error)); ok {
		return rf(ctx, customerID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*resources.Resource); ok {
		r0 = rf(ctx, customerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*resources.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, customerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateResource provides a mock function with given fields: ctx, resource
func (_m *MockRepository) UpdateResource(ctx context.Context, resource *resources.Resource) error {
	ret := _m.Called(ctx, resource)

	if len(ret) == 0 {
		panic("no return value specified for UpdateResource")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *resources.Resource) error); ok {
		r0 = rf(ctx, resource)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockRepository creates a new instance of MockRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRepository {
	mock := &MockRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}