// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	abstraction "moonlay-todolist/internal/abstraction"

	mock "github.com/stretchr/testify/mock"

	model "moonlay-todolist/internal/model"
)

// IListRepository is an autogenerated mock type for the IListRepository type
type IListRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *IListRepository) Create(data *model.List) (*model.List, error) {
	ret := _m.Called(data)

	var r0 *model.List
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.List) (*model.List, error)); ok {
		return rf(data)
	}
	if rf, ok := ret.Get(0).(func(*model.List) *model.List); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.List)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.List) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateFile provides a mock function with given fields: link, listID
func (_m *IListRepository) CreateFile(link string, listID string) (*model.ListFile, error) {
	ret := _m.Called(link, listID)

	var r0 *model.ListFile
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (*model.ListFile, error)); ok {
		return rf(link, listID)
	}
	if rf, ok := ret.Get(0).(func(string, string) *model.ListFile); ok {
		r0 = rf(link, listID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ListFile)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(link, listID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteByID provides a mock function with given fields: ID
func (_m *IListRepository) DeleteByID(ID string) (*model.List, error) {
	ret := _m.Called(ID)

	var r0 *model.List
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.List, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.List); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.List)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteFileByListID provides a mock function with given fields: ListID
func (_m *IListRepository) DeleteFileByListID(ListID string) error {
	ret := _m.Called(ListID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(ListID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: s, d, t, p
func (_m *IListRepository) FindAll(s *bool, d *string, t *string, p *abstraction.Pagination) (*[]model.List, *abstraction.PaginationInfo, error) {
	ret := _m.Called(s, d, t, p)

	var r0 *[]model.List
	var r1 *abstraction.PaginationInfo
	var r2 error
	if rf, ok := ret.Get(0).(func(*bool, *string, *string, *abstraction.Pagination) (*[]model.List, *abstraction.PaginationInfo, error)); ok {
		return rf(s, d, t, p)
	}
	if rf, ok := ret.Get(0).(func(*bool, *string, *string, *abstraction.Pagination) *[]model.List); ok {
		r0 = rf(s, d, t, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]model.List)
		}
	}

	if rf, ok := ret.Get(1).(func(*bool, *string, *string, *abstraction.Pagination) *abstraction.PaginationInfo); ok {
		r1 = rf(s, d, t, p)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*abstraction.PaginationInfo)
		}
	}

	if rf, ok := ret.Get(2).(func(*bool, *string, *string, *abstraction.Pagination) error); ok {
		r2 = rf(s, d, t, p)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FindByID provides a mock function with given fields: ID
func (_m *IListRepository) FindByID(ID string) (*model.List, error) {
	ret := _m.Called(ID)

	var r0 *model.List
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.List, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(string) *model.List); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.List)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateByID provides a mock function with given fields: ID, data
func (_m *IListRepository) UpdateByID(ID string, data *model.List) (*model.List, error) {
	ret := _m.Called(ID, data)

	var r0 *model.List
	var r1 error
	if rf, ok := ret.Get(0).(func(string, *model.List) (*model.List, error)); ok {
		return rf(ID, data)
	}
	if rf, ok := ret.Get(0).(func(string, *model.List) *model.List); ok {
		r0 = rf(ID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.List)
		}
	}

	if rf, ok := ret.Get(1).(func(string, *model.List) error); ok {
		r1 = rf(ID, data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIListRepository creates a new instance of IListRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIListRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IListRepository {
	mock := &IListRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
