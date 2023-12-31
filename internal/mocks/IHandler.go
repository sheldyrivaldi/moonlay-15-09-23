// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// IHandler is an autogenerated mock type for the IHandler type
type IHandler struct {
	mock.Mock
}

// Create provides a mock function with given fields: c
func (_m *IHandler) Create(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: c
func (_m *IHandler) DeleteByID(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: c
func (_m *IHandler) FindAll(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByID provides a mock function with given fields: c
func (_m *IHandler) FindByID(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateByID provides a mock function with given fields: c
func (_m *IHandler) UpdateByID(c echo.Context) error {
	ret := _m.Called(c)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIHandler creates a new instance of IHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *IHandler {
	mock := &IHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
