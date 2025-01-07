package No_63

import (
	"fmt"
	"testing"
	"time"
)

func TestForGoroutine(t *testing.T) {
	s := []int{1, 2, 3}
	for _, v := range s {
		go func(val int) {
			fmt.Println(val)
		}(v)
	}
	time.Sleep(2 * time.Second)
}
