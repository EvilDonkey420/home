// Code generated by mockery v1.0.0
package mocks

import mock "github.com/stretchr/testify/mock"

// Subscription is an autogenerated mock type for the Subscription type
type Subscription struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Subscription) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Events provides a mock function with given fields:
func (_m *Subscription) Events() <-chan interface{} {
	ret := _m.Called()

	var r0 <-chan interface{}
	if rf, ok := ret.Get(0).(func() <-chan interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan interface{})
		}
	}

	return r0
}

// ID provides a mock function with given fields:
func (_m *Subscription) ID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Write provides a mock function with given fields: event
func (_m *Subscription) Write(event interface{}) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
