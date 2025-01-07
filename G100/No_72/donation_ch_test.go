package No_72

import (
	"fmt"
	"testing"
	"time"
)

func TestChDon(t *testing.T) {
	donation1 := Donation1{
		ch: make(chan int),
	}
	f := func(goal int) {
		for c := range donation1.ch {
			if c > goal {
				fmt.Printf("$%d goal reacher\n", goal)
			}
		}
	}

	go f(10)
	go f(15)

	go func() {
		for {
			time.Sleep(time.Second)
			donation1.balance++
			donation1.ch <- donation1.balance
		}
	}()
}

type Donation1 struct {
	balance int
	ch      chan int
}
