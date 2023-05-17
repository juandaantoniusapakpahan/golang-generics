package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GG int

type NumberA interface {
	~int | int32 | int64 | float32 | float64
}

func Mins[T SetNumber](param1, param2 T) T {
	if param1 < param2 {
		return param1
	} else if param1 > param2 {
		return param2
	}
	return 0
}

func TestInferenc(t *testing.T) {
	assert.Equal(t, 100, Mins(100, 200))             // inference
	assert.Equal(t, 100, Mins(100, 200))             // inference
	assert.Equal(t, GG(100), Mins(GG(100), GG(200))) // inference

}

// go test -v -run=TestInferenc
