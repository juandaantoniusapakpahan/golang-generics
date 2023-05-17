package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](param1, param2 T) bool {
	if param1 == param2 {
		return true
	}
	return false
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[int](100, 100))
	assert.True(t, IsSame[string]("joke", "joke"))
	assert.True(t, IsSame[bool](true, true))
	assert.False(t, IsSame[int](102, 100))
	assert.False(t, IsSame[string]("joke", "joge"))

}

// go test -v -run=TestIsSame
