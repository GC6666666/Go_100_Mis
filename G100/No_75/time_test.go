package No_75

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ticker := time.NewTicker(1 * time.Millisecond)
	for i := 0; i < 10; i++ {
		select {
		case <-ticker.C:
			fmt.Println(i)
		}
	}
}
