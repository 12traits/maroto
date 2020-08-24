// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import consts "github.com/12traits/maroto/pkg/consts"

import mock "github.com/stretchr/testify/mock"

// Font is an autogenerated mock type for the Font type
type Font struct {
	mock.Mock
}

// GetFamily provides a mock function with given fields:
func (_m *Font) GetFamily() consts.Family {
	ret := _m.Called()

	var r0 consts.Family
	if rf, ok := ret.Get(0).(func() consts.Family); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(consts.Family)
	}

	return r0
}

// GetFont provides a mock function with given fields:
func (_m *Font) GetFont() (consts.Family, consts.Style, float64) {
	ret := _m.Called()

	var r0 consts.Family
	if rf, ok := ret.Get(0).(func() consts.Family); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(consts.Family)
	}

	var r1 consts.Style
	if rf, ok := ret.Get(1).(func() consts.Style); ok {
		r1 = rf()
	} else {
		r1 = ret.Get(1).(consts.Style)
	}

	var r2 float64
	if rf, ok := ret.Get(2).(func() float64); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(float64)
	}

	return r0, r1, r2
}

// GetScaleFactor provides a mock function with given fields:
func (_m *Font) GetScaleFactor() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetSize provides a mock function with given fields:
func (_m *Font) GetSize() float64 {
	ret := _m.Called()

	var r0 float64
	if rf, ok := ret.Get(0).(func() float64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetStyle provides a mock function with given fields:
func (_m *Font) GetStyle() consts.Style {
	ret := _m.Called()

	var r0 consts.Style
	if rf, ok := ret.Get(0).(func() consts.Style); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(consts.Style)
	}

	return r0
}

// SetFamily provides a mock function with given fields: family
func (_m *Font) SetFamily(family consts.Family) {
	_m.Called(family)
}

// SetFont provides a mock function with given fields: family, style, size
func (_m *Font) SetFont(family consts.Family, style consts.Style, size float64) {
	_m.Called(family, style, size)
}

// SetSize provides a mock function with given fields: size
func (_m *Font) SetSize(size float64) {
	_m.Called(size)
}

// SetStyle provides a mock function with given fields: style
func (_m *Font) SetStyle(style consts.Style) {
	_m.Called(style)
}
