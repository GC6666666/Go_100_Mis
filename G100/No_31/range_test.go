package No_31

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	n := [3]int{1, 2, 3}  // &n = 0x001
	for i, v := range n { // p = *[3]int = 0x001
		n[1] = 20000
		n[2] = 10000
		if i == 1 {
			fmt.Println(v)
		}
		if i == 2 {
			fmt.Println(v)
		}
	}
}
