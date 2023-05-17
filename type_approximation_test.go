package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Age int

type SetNumber interface {
	~int | int32 | int64 | float32 | float64
}

func Min[T SetNumber](param1, param2 T) T {
	if param1 < param2 {
		return param1
	} else if param1 > param2 {
		return param2
	}
	return 0
}

func TestApproximation(t *testing.T) {
	assert.Equal(t, Age(100), Min[Age](Age(100), Age(200))) // Approximation
	assert.Equal(t, float32(100), Min[float32](float32(100), float32(200)))

}

// go test -v -run=TestApproximation
