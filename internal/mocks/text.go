// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	internal "github.com/johnfercher/maroto/internal"
	mock "github.com/stretchr/testify/mock"

	props "github.com/johnfercher/maroto/pkg/props"
)

// Text is an autogenerated mock type for the Text type
type Text struct {
	mock.Mock
}

// Add provides a mock function with given fields: text, cell, textProp
func (_m *Text) Add(text string, cell internal.Cell, textProp props.Text) {
	_m.Called(text, cell, textProp)
}

// GetLinesQuantity provides a mock function with given fields: text, fontFamily, colWidth
func (_m *Text) GetLinesQuantity(text string, fontFamily props.Text, colWidth float64) int {
	ret := _m.Called(text, fontFamily, colWidth)

	var r0 int
	if rf, ok := ret.Get(0).(func(string, props.Text, float64) int); ok {
		r0 = rf(text, fontFamily, colWidth)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}
