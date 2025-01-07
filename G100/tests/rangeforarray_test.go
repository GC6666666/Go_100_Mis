package tests

import (
	"fmt"
	"testing"
)

func TestForArray(t *testing.T) {
	n := [3]int{1, 2, 3}
	for i, v := range &n {
		n[2] = 10
		if i == 2 {
			fmt.Println(v)
		}
	}
}
