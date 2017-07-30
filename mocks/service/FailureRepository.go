// Code generated by mockery v1.0.0
package service

import failure "github.com/slok/ragnarok/failure"
import mock "github.com/stretchr/testify/mock"

// FailureRepository is an autogenerated mock type for the FailureRepository type
type FailureRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id
func (_m *FailureRepository) Delete(id string) {
	_m.Called(id)
}

// Get provides a mock function with given fields: id
func (_m *FailureRepository) Get(id string) (*failure.Failure, bool) {
	ret := _m.Called(id)

	var r0 *failure.Failure
	if rf, ok := ret.Get(0).(func(string) *failure.Failure); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*failure.Failure)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields:
func (_m *FailureRepository) GetAll() []*failure.Failure {
	ret := _m.Called()

	var r0 []*failure.Failure
	if rf, ok := ret.Get(0).(func() []*failure.Failure); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*failure.Failure)
		}
	}

	return r0
}

// GetAllByNode provides a mock function with given fields: nodeID
func (_m *FailureRepository) GetAllByNode(nodeID string) []*failure.Failure {
	ret := _m.Called(nodeID)

	var r0 []*failure.Failure
	if rf, ok := ret.Get(0).(func(string) []*failure.Failure); ok {
		r0 = rf(nodeID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*failure.Failure)
		}
	}

	return r0
}

// Store provides a mock function with given fields: _a0
func (_m *FailureRepository) Store(_a0 *failure.Failure) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*failure.Failure) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
