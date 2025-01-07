package tests

import "testing"

func TestRangeFor(t *testing.T) {
	s := []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		s = append(s, 10)
	}

}
