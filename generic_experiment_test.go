package golanggenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, int(100), ExperimentalMin[int](int(100), int(200)))
	assert.Equal(t, int32(100), ExperimentalMin[int32](int32(100), int32(200)))
}

func TestMap(t *testing.T) {
	data1 := map[string]string{
		"name": "GGWP",
	}
	data2 := map[string]string{
		"name": "GGWP",
	}

	assert.True(t, maps.Equal(data1, data2))
}
func TestSlice(t *testing.T) {
	data1 := []int{120, 3023}
	data2 := []int{120, 3023}

	data3 := []string{"gojek"}
	data4 := []string{"gojek"}

	assert.True(t, slices.Equal(data1, data2))
	assert.True(t, slices.Equal(data3, data4))

}
