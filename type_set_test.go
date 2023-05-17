package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int32 | int64 | float32 | float64
}

func MinNumber[T Number](number1 T, number2 T) T {
	if number1 < number2 {
		return number1
	} else if number2 < number1 {
		return number2
	}
	return 0
}
func TestSetType(t *testing.T) {
	assert.Equal(t, 5, MinNumber[int](10, 5))
	assert.Equal(t, 3.1, MinNumber[float64](3.1, 5))
	assert.Equal(t, int32(3), MinNumber[int32](3, 5))

}

// go test -v -run=TestSetType
