package tests

import (
	"fmt"
	"testing"
)

func TestForSwitchBreak(t *testing.T) {
loop:
	for i := 0; i < 5; i++ {
		fmt.Printf("%d", i)
		switch i {
		case 2:
			break loop
		default:

		}
	}

}
