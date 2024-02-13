// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Type is an autogenerated mock type for the Type type
type Type struct {
	mock.Mock
}

// Match provides a mock function with given fields: int0
func (_m *Type) Match(int0 int) bool {
	ret := _m.Called(int0)

	if len(ret) == 0 {
		panic("no return value specified for Match")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(int) bool); ok {
		r0 = rf(int0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Name provides a mock function with given fields:
func (_m *Type) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewType creates a new instance of Type. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewType(t interface {
	mock.TestingT
	Cleanup(func())
}) *Type {
	mock := &Type{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
