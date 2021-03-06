// Code generated by mockery v1.0.0
package clock

import mock "github.com/stretchr/testify/mock"
import time "time"

// Clock is an autogenerated mock type for the Clock type
type Clock struct {
	mock.Mock
}

// After provides a mock function with given fields: d
func (_m *Clock) After(d time.Duration) <-chan time.Time {
	ret := _m.Called(d)

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func(time.Duration) <-chan time.Time); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}

// NewTicker provides a mock function with given fields: d
func (_m *Clock) NewTicker(d time.Duration) *time.Ticker {
	ret := _m.Called(d)

	var r0 *time.Ticker
	if rf, ok := ret.Get(0).(func(time.Duration) *time.Ticker); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Ticker)
		}
	}

	return r0
}

// NewTimer provides a mock function with given fields: d
func (_m *Clock) NewTimer(d time.Duration) *time.Timer {
	ret := _m.Called(d)

	var r0 *time.Timer
	if rf, ok := ret.Get(0).(func(time.Duration) *time.Timer); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Timer)
		}
	}

	return r0
}

// Now provides a mock function with given fields:
func (_m *Clock) Now() time.Time {
	ret := _m.Called()

	var r0 time.Time
	if rf, ok := ret.Get(0).(func() time.Time); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(time.Time)
	}

	return r0
}

// Sleep provides a mock function with given fields: d
func (_m *Clock) Sleep(d time.Duration) {
	_m.Called(d)
}

// Tick provides a mock function with given fields: d
func (_m *Clock) Tick(d time.Duration) <-chan time.Time {
	ret := _m.Called(d)

	var r0 <-chan time.Time
	if rf, ok := ret.Get(0).(func(time.Duration) <-chan time.Time); ok {
		r0 = rf(d)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan time.Time)
		}
	}

	return r0
}
