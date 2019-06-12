// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

import v1 "github.com/dairlair/twitwatch/pkg/api/v1"

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// AddStream provides a mock function with given fields: stream
func (_m *Interface) AddStream(stream *v1.Stream) (int64, error) {
	ret := _m.Called(stream)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*v1.Stream) int64); ok {
		r0 = rf(stream)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Stream) error); ok {
		r1 = rf(stream)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddTwit provides a mock function with given fields: twit
func (_m *Interface) AddTwit(twit *v1.Twit) (int64, error) {
	ret := _m.Called(twit)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*v1.Twit) int64); ok {
		r0 = rf(twit)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Twit) error); ok {
		r1 = rf(twit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
