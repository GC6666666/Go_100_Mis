package tests

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringIteration(t *testing.T) {
	//s := "hêllo"
	//for i := range s {
	//	fmt.Printf("%d location is %c\n", i, s[i])
	//}
	//fmt.Println("len of s is", len(s))
	//fmt.Println("len of rune(s) is", utf8.RuneCountInString(s))
	//for i, v := range s {
	//	fmt.Printf("%d location is %c\n", i, v)
	//}
	//
	//ss := []rune(s)
	//for i := range ss {
	//	fmt.Printf("%d location is %c\n", i, ss[i])
	//}
	s := "hêllo"
	char := []rune(s)[1]
	fmt.Printf("char is %c\n", char)
	fmt.Println(strings.TrimSuffix("123oxo", "xo"))
}
