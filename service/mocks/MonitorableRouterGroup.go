// Code generated by mockery v2.46.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"

	options "github.com/monitoror/monitoror/service/options"
)

// MonitorableRouterGroup is an autogenerated mock type for the MonitorableRouterGroup type
type MonitorableRouterGroup struct {
	mock.Mock
}

// GET provides a mock function with given fields: path, handlerFunc, _a2
func (_m *MonitorableRouterGroup) GET(path string, handlerFunc echo.HandlerFunc, _a2 ...options.RouterOption) *echo.Route {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, path, handlerFunc)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GET")
	}

	var r0 *echo.Route
	if rf, ok := ret.Get(0).(func(string, echo.HandlerFunc, ...options.RouterOption) *echo.Route); ok {
		r0 = rf(path, handlerFunc, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*echo.Route)
		}
	}

	return r0
}

// NewMonitorableRouterGroup creates a new instance of MonitorableRouterGroup. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMonitorableRouterGroup(t interface {
	mock.TestingT
	Cleanup(func())
}) *MonitorableRouterGroup {
	mock := &MonitorableRouterGroup{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
