package golanggenerics

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	Name  T
	Email T
}

func (data *Data[_]) GetName(name string) string {
	return "Hello " + name
}

func (data *Data[T]) GetEmail(email T) T {
	data.Email = email
	return data.Email
}

func TestStruct(t *testing.T) {
	data := Data[string]{
		Name:  "GGWP",
		Email: "GGEWp",
	}

	fmt.Println(data)
}

func TestStructMethod(t *testing.T) {
	data := &Data[string]{
		Name:  "GGWP",
		Email: "GGEWp",
	}

	fmt.Println(data.GetName("ggwp"))
	fmt.Println(data.GetEmail("ggwp@gmail.com"))
}
