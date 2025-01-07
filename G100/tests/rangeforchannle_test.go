package tests

import (
	"fmt"
	"testing"
)

func TestRangeForChannel(t *testing.T) {
	ch1 := make(chan int, 3)
	go func() {
		ch1 <- 11
		ch1 <- 12
		ch1 <- 13
		// 不要忘记关闭
		close(ch1)
	}()

	ch2 := make(chan int, 3)
	go func() {
		ch2 <- 21
		ch2 <- 22
		ch2 <- 23
		close(ch2)
	}()

	ch := ch1
	for c := range ch {
		fmt.Println(c)
		ch = ch2
	}
	fmt.Println()
	for cc := range ch {
		fmt.Println(cc)
	}
}
