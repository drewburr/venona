// Copyright 2020 The Codefresh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v1.1.1. DO NOT EDIT.

package mocks

import (
	log15 "github.com/inconshreveable/log15"

	mock "github.com/stretchr/testify/mock"
)

// Logger is an autogenerated mock type for the Logger type
type Logger struct {
	mock.Mock
}

// Crit provides a mock function with given fields: msg, ctx
func (_m *Logger) Crit(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}

// Debug provides a mock function with given fields: msg, ctx
func (_m *Logger) Debug(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: msg, ctx
func (_m *Logger) Error(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}

// GetHandler provides a mock function with given fields:
func (_m *Logger) GetHandler() log15.Handler {
	ret := _m.Called()

	var r0 log15.Handler
	if rf, ok := ret.Get(0).(func() log15.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log15.Handler)
		}
	}

	return r0
}

// Info provides a mock function with given fields: msg, ctx
func (_m *Logger) Info(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}

// New provides a mock function with given fields: ctx
func (_m *Logger) New(ctx ...interface{}) log15.Logger {
	var _ca []interface{}
	_ca = append(_ca, ctx...)
	ret := _m.Called(_ca...)

	var r0 log15.Logger
	if rf, ok := ret.Get(0).(func(...interface{}) log15.Logger); ok {
		r0 = rf(ctx...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(log15.Logger)
		}
	}

	return r0
}

// SetHandler provides a mock function with given fields: h
func (_m *Logger) SetHandler(h log15.Handler) {
	_m.Called(h)
}

// Warn provides a mock function with given fields: msg, ctx
func (_m *Logger) Warn(msg string, ctx ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, ctx...)
	_m.Called(_ca...)
}
