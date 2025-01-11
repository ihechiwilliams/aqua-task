// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	customers "aqua-backend/internal/repositories/customers"
	context "context"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockRepository is an autogenerated mock type for the Repository type
type MockRepository struct {
	mock.Mock
}

// CreateCustomer provides a mock function with given fields: ctx, customer
func (_m *MockRepository) CreateCustomer(ctx context.Context, customer *customers.DBCustomer) (*customers.Customer, error) {
	ret := _m.Called(ctx, customer)

	if len(ret) == 0 {
		panic("no return value specified for CreateCustomer")
	}

	var r0 *customers.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *customers.DBCustomer) (*customers.Customer, error)); ok {
		return rf(ctx, customer)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *customers.DBCustomer) *customers.Customer); ok {
		r0 = rf(ctx, customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customers.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *customers.DBCustomer) error); ok {
		r1 = rf(ctx, customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCustomerByID provides a mock function with given fields: ctx, id
func (_m *MockRepository) GetCustomerByID(ctx context.Context, id uuid.UUID) (*customers.Customer, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetCustomerByID")
	}

	var r0 *customers.Customer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*customers.Customer, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *customers.Customer); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customers.Customer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
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
