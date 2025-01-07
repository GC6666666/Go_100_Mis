package tests

import (
	"fmt"
	"testing"
)

func TestClosure(t *testing.T) {
	//i, j := 0, 0
	//defer func(i int) {
	//	fmt.Println(i, j)
	//}(i)
	//i++
	//j++
	s := Struct{
		id: "foo",
	}
	//defer s.print()
	s.id = "bar"
	s.print()
}

type Struct struct {
	id string
}

func (s Struct) print() {
	fmt.Println(s.id)
}
