package golanggenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// any = interface{}
func GenericShow[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestType(t *testing.T) {
	var intValue int = GenericShow[int](1000)
	assert.Equal(t, 1000, intValue)

	var strValue string = GenericShow[string]("Hello World")
	assert.Equal(t, "Hello World", strValue)

	var fltValue float32 = GenericShow[float32](10.22)
	assert.Equal(t, float32(10.22), fltValue)
}
