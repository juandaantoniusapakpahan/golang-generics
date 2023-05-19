package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func MinValue[T interface{ int | int32 | int64 | float64 }](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestInline(t *testing.T) {
	assert.Equal(t, int32(100), MinValue[int32](int32(100), int32(200)))
	assert.Equal(t, int64(100), MinValue[int64](int64(100), int64(200)))
}

func GetFirst[T []E, E any](slice T) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {

	data := []int{100, 2230, 23, 232}
	data2 := []string{"babu", "check", "out"}
	assert.Equal(t, 100, GetFirst[[]int, int](data))
	assert.Equal(t, "babu", GetFirst[[]string, string](data2))

}
