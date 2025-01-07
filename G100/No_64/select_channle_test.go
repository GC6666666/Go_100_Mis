package No_64

import (
	"fmt"
	"testing"
)

func TestForGoroutine(t *testing.T) {
	messagech := make(chan int)
	disconnectch := make(chan struct{})
	n := 10
	go func() {
		for i := 0; i < n; i++ {
			messagech <- i
		}
		disconnectch <- struct{}{}
	}()

	for {
		select {
		case val := <-messagech:
			fmt.Println(val)
		case <-disconnectch:
			fmt.Println("disconnection, return")
			return
		}
	}

}

func merge(ch1, ch2 <-chan int) chan int {
	ch := make(chan int)

	go func() {
		for {
			select {
			case v := <-ch1:
				ch <- v
			case v := <-ch2:
				ch <- v
			}
		}
		close(ch)
	}()
	return ch
}
