package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (data *MyData[T]) SetValue(number T) {
	data.Value = number
}

func (data *MyData[T]) GetValue() T {
	return data.Value
}

func TestInterfaceGeneric(t *testing.T) {
	mydata := MyData[string]{}
	result := ChangeValue[string](&mydata, "GWP")

	assert.Equal(t, "GWP", result)
	assert.Equal(t, "GWP", mydata.Value)
}
