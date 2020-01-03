// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Watcher is an autogenerated mock type for the Watcher type
type Watcher struct {
	mock.Mock
}

// AddFunc provides a mock function with given fields: obj
func (_m *Watcher) AddFunc(obj interface{}) {
	_m.Called(obj)
}

// Run provides a mock function with given fields:
func (_m *Watcher) Run() {
	_m.Called()
}

// UpdateFunc provides a mock function with given fields: oldobj, newobj
func (_m *Watcher) UpdateFunc(oldobj interface{}, newobj interface{}) {
	_m.Called(oldobj, newobj)
}
