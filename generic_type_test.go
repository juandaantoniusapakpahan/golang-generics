package golanggenerics

import (
	"fmt"
	"testing"
)

type Box[T any] []T

// type M map[string]interface{}

func Count[T any](box Box[T]) {

	for _, value := range box {
		fmt.Println(value)
	}
}

func TestGenericType(t *testing.T) {
	Count[int](Box[int]{1, 23, 4, 5, 5, 6})
	Count[string](Box[string]{"1", "2", "3", "4", "45"})

}

// go test -v -run=TestGenericType
