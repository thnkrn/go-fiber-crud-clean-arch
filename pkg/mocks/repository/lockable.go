// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Lockable is an autogenerated mock type for the Lockable type
type Lockable struct {
	mock.Mock
}

// GetVersion provides a mock function with given fields:
func (_m *Lockable) GetVersion() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// SetVersion provides a mock function with given fields: version
func (_m *Lockable) SetVersion(version int) {
	_m.Called(version)
}

// NewLockable creates a new instance of Lockable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLockable(t interface {
	mock.TestingT
	Cleanup(func())
}) *Lockable {
	mock := &Lockable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}