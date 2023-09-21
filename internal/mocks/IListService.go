// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	listdto "moonlay-todolist/internal/dto/list"
)

// ISubListService is an autogenerated mock type for the ISubListService type
type IListService struct {
	mock.Mock
}

// Create provides a mock function with given fields: payload
func (_m *IListService) Create(payload *listdto.CreateRequest) (*listdto.CreateResponse, error) {
	ret := _m.Called(payload)

	var r0 *listdto.CreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*listdto.CreateRequest) (*listdto.CreateResponse, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*listdto.CreateRequest) *listdto.CreateResponse); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listdto.CreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*listdto.CreateRequest) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: payload
func (_m *IListService) DeleteByID(payload *listdto.DeleteByIDRequest) (*listdto.DeleteByIDResponse, error) {
	ret := _m.Called(payload)

	var r0 *listdto.DeleteByIDResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*listdto.DeleteByIDRequest) (*listdto.DeleteByIDResponse, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*listdto.DeleteByIDRequest) *listdto.DeleteByIDResponse); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listdto.DeleteByIDResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*listdto.DeleteByIDRequest) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: payload
func (_m *IListService) FindAll(payload *listdto.FindAllRequest) (*listdto.FindAllResponse, error) {
	ret := _m.Called(payload)

	var r0 *listdto.FindAllResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*listdto.FindAllRequest) (*listdto.FindAllResponse, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*listdto.FindAllRequest) *listdto.FindAllResponse); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listdto.FindAllResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*listdto.FindAllRequest) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: payload
func (_m *IListService) FindByID(payload *listdto.FindByIDRequest) (*listdto.FindByIDResponse, error) {
	ret := _m.Called(payload)

	var r0 *listdto.FindByIDResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*listdto.FindByIDRequest) (*listdto.FindByIDResponse, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*listdto.FindByIDRequest) *listdto.FindByIDResponse); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listdto.FindByIDResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*listdto.FindByIDRequest) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: payload, files
func (_m *IListService) UpdateByID(payload *listdto.UpdateByIDRequest, files []string) (*listdto.UpdateByIDResponse, error) {
	ret := _m.Called(payload, files)

	var r0 *listdto.UpdateByIDResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*listdto.UpdateByIDRequest, []string) (*listdto.UpdateByIDResponse, error)); ok {
		return rf(payload, files)
	}
	if rf, ok := ret.Get(0).(func(*listdto.UpdateByIDRequest, []string) *listdto.UpdateByIDResponse); ok {
		r0 = rf(payload, files)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*listdto.UpdateByIDResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*listdto.UpdateByIDRequest, []string) error); ok {
		r1 = rf(payload, files)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIListService creates a new instance of IListService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIListService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IListService {
	mock := &IListService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}