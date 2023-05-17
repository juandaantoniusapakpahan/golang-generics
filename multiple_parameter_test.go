package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// any = interface{}
func MultipleParamaGeneric[T any, T2 any](param1 T, param2 T2) (T, T2) {
	fmt.Println(param1)
	fmt.Println(param2)
	return param1, param2
}

func TestMultipleParamaGeneric(t *testing.T) {
	param1, param2 := MultipleParamaGeneric[int, string](100, "Hello String")

	assert.Equal(t, 100, param1)
	assert.Equal(t, "Hello String", param2)
}

// go test -v -run=TestMultipleParamaGeneric
