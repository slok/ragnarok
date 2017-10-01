// Code generated by mockery v1.0.0
package service

import failure "github.com/slok/ragnarok/failure"
import mock "github.com/stretchr/testify/mock"

// FailureState is an autogenerated mock type for the FailureState type
type FailureState struct {
	mock.Mock
}

// ProcessFailureStates provides a mock function with given fields: failures
func (_m *FailureState) ProcessFailureStates(failures []*failure.Failure) error {
	ret := _m.Called(failures)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*failure.Failure) error); ok {
		r0 = rf(failures)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StartHandling provides a mock function with given fields:
func (_m *FailureState) StartHandling() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StopHandling provides a mock function with given fields:
func (_m *FailureState) StopHandling() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}