package tests

import (
	"fmt"
	"testing"
)

func TestRune(t *testing.T) {
	s := "hello,世界"
	fmt.Println(s[7])
	for index, char := range s {
		fmt.Printf("字符 %c 在位置 %d\n", char, index)
	}
}
