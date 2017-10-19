// Code generated by mockery v1.0.0
package client

import mock "github.com/stretchr/testify/mock"
import v1 "github.com/slok/ragnarok/api/cluster/v1"

// Status is an autogenerated mock type for the Status type
type Status struct {
	mock.Mock
}

// NodeHeartbeat provides a mock function with given fields: id, status
func (_m *Status) NodeHeartbeat(id string, status v1.NodeState) error {
	ret := _m.Called(id, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, v1.NodeState) error); ok {
		r0 = rf(id, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterNode provides a mock function with given fields: id, tags
func (_m *Status) RegisterNode(id string, tags map[string]string) error {
	ret := _m.Called(id, tags)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, map[string]string) error); ok {
		r0 = rf(id, tags)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
