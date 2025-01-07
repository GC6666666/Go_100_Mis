package tests

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestEmptyOrNil(t *testing.T) {
	//var s []string
	//loc(1, s)
	//
	//s = []string(nil)
	//loc(2, s)
	//
	//s = []string{}
	//loc(3, s)
	//
	//s = make([]string, 0)
	//loc(4, s)
	var s1 []float32
	customer1 := myCustomer{
		ID:     "foo",
		Option: s1,
	}
	b, _ := json.Marshal(customer1)
	fmt.Println(string(b))

	s2 := make([]float32, 0)
	customer2 := myCustomer{
		ID:     "foo",
		Option: s2,
	}
	c, _ := json.Marshal(customer2)
	fmt.Println(string(c))
}

type myCustomer struct {
	ID     string
	Option []float32
}

func llog(i int, s []string) {
	fmt.Printf("%d : empty = %t\tnil = %t\n", i, len(s) == 0, s == nil)
	fmt.Println(s)
}
