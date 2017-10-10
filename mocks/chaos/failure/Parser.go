// Code generated by mockery v1.0.0
package failure

import failure "github.com/slok/ragnarok/chaos/failure"
import failurestatus "github.com/slok/ragnarok/grpc/failurestatus"
import mock "github.com/stretchr/testify/mock"

// Parser is an autogenerated mock type for the Parser type
type Parser struct {
	mock.Mock
}

// FailureToPB provides a mock function with given fields: fl
func (_m *Parser) FailureToPB(fl *failure.Failure) (*failurestatus.Failure, error) {
	ret := _m.Called(fl)

	var r0 *failurestatus.Failure
	if rf, ok := ret.Get(0).(func(*failure.Failure) *failurestatus.Failure); ok {
		r0 = rf(fl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*failurestatus.Failure)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failure.Failure) error); ok {
		r1 = rf(fl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PBToFailure provides a mock function with given fields: fl
func (_m *Parser) PBToFailure(fl *failurestatus.Failure) (*failure.Failure, error) {
	ret := _m.Called(fl)

	var r0 *failure.Failure
	if rf, ok := ret.Get(0).(func(*failurestatus.Failure) *failure.Failure); ok {
		r0 = rf(fl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*failure.Failure)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*failurestatus.Failure) error); ok {
		r1 = rf(fl)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}