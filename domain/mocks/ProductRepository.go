// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	"context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/zeintkp/go-rest/domain"
)

// ProductRepository is an autogenerated mock type for the ProductRepository type
type ProductRepository struct {
	mock.Mock
}

// Browse provides a mock function with given fields: search, order, sort, limit, offset
func (_m *ProductRepository) Browse(ctx context.Context, search, order, sort string, limit, offset int) ([]domain.Product, int, error) {
	ret := _m.Called(ctx, search, order, sort, limit, offset)

	var r0 []domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, int, int) []domain.Product); ok {
		r0 = rf(ctx, search, order, sort, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Product)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, int, int) int); ok {
		r1 = rf(ctx, search, order, sort, limit, offset)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, string, int, int) error); ok {
		r2 = rf(ctx, search, order, sort, limit, offset)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Create provides a mock function with given fields: a1
func (_m *ProductRepository) Create(ctx context.Context, a1 domain.Product) error {
	ret := _m.Called(ctx, a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) error); ok {
		r0 = rf(ctx, a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Read provides a mock function with given fields: id
func (_m *ProductRepository) Read(ctx context.Context, id string) (domain.Product, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.Product
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Product); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(domain.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ar
func (_m *ProductRepository) Update(ctx context.Context, ar domain.Product) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ar
func (_m *ProductRepository) Delete(ctx context.Context, ar domain.Product) error {
	ret := _m.Called(ctx, ar)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.Product) error); ok {
		r0 = rf(ctx, ar)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Read provides a mock function with given fields: id
func (_m *ProductRepository) IsExist(ctx context.Context, id, name string) (bool, error) {
	ret := _m.Called(ctx, id, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, id, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, id, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}