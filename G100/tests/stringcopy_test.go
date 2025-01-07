package tests

import (
	"fmt"
	"testing"
)

func TestStringCopy(t *testing.T) {
	s := "hêllo,world!"
	fmt.Println(s[:5])
	fmt.Printf("%c\n", []rune(s)[:5])
}
