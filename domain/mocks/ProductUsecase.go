// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import context "context"
import domain "github.com/zeintkp/go-rest/domain"
import mock "github.com/stretchr/testify/mock"

// ProductUsecase is an autogenerated mock type for the ProductUsecase type
type ProductUsecase struct {
	mock.Mock
}

// Browse provides a mock function with given fields: ctx, cursor, num
func (_m *ProductUsecase) Browse(ctx context.Context, search, order, sort string, limit, page int) ([]domain.ProductVM, domain.PaginationVM, error) {
	ret := _m.Called(ctx, search, order, sort, limit, page)

	var r0 []domain.ProductVM
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, int, int) []domain.ProductVM); ok {
		r0 = rf(ctx, search, order, sort, limit, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.ProductVM)
		}
	}

	var r1 domain.PaginationVM
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, int, int) domain.PaginationVM); ok {
		r1 = rf(ctx, search, order, sort, limit, page)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(domain.PaginationVM)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, string, string, int, int) error); ok {
		r2 = rf(ctx, search, order, sort, limit, page)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Create provides a mock function with given fields: ar
func (_m *ProductUsecase) Create(ctx context.Context, ar domain.ProductRequest) (domain.ProductVM, error) {
	ret := _m.Called(ctx, ar)

	var r0 domain.ProductVM
	if rf, ok := ret.Get(0).(func(context.Context, domain.ProductRequest) domain.ProductVM); ok {
		r0 = rf(ctx, ar)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.ProductVM)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.ProductRequest) error); ok {
		r1 = rf(ctx, ar)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: id
func (_m *ProductUsecase) Read(ctx context.Context, id string) (domain.ProductVM, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.ProductVM
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ProductVM); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.ProductVM)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, ar
func (_m *ProductUsecase) Update(ctx context.Context, ar domain.ProductRequest, id string) (domain.ProductVM, error) {
	ret := _m.Called(ctx, ar, id)

	var r0 domain.ProductVM
	if rf, ok := ret.Get(0).(func(context.Context, domain.ProductRequest, string) domain.ProductVM); ok {
		r0 = rf(ctx, ar, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.ProductVM)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.ProductRequest, string) error); ok {
		r1 = rf(ctx, ar, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ProductUsecase) Delete(ctx context.Context, id string) (domain.ProductVM, error) {
	ret := _m.Called(ctx, id)

	var r0 domain.ProductVM
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.ProductVM); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.ProductVM)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Create provides a mock function with given fields: ar
func (_m *ProductUsecase) ConvertVM(ar domain.Product) domain.ProductVM {
	ret := _m.Called(ar)

	var r0 domain.ProductVM
	if rf, ok := ret.Get(0).(func(domain.Product) domain.ProductVM); ok {
		r0 = rf(ar)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(domain.ProductVM)
		}
	}

	return r0
}
