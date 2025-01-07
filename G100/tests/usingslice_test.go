package tests

import (
	"fmt"
	"testing"
)

func TestUsingSlice(t *testing.T) {
	//s := make([]int, 3, 6)
	////msl.S[4] = 1
	//s[1] = 1
	//s = append(s, 2)
	//s = append(s, 3, 4, 5)
	//fmt.Println(s)
	//fmt.Println(cap(s))
	//s1 := make([]int, 3, 6)
	//	//s2 := s1[1:3]
	//	//s1[1] = 1
	//	//fmt.Println(s1)
	//	//fmt.Println(s2)
	//	//s2 = append(s2, 2)
	//	//s2 = append(s2, 3, 4, 5)
	//	//fmt.Println("s1:", s1)
	//	//fmt.Println("s2:", s2)
	//	//s1 = append(s1, 6)
	//	//fmt.Println("s1:", s1)
	//	//fmt.Println("s2:", s2)
	s1 := []int{1, 2, 3}
	s2 := s1[:2:2]
	s3 := append(s2, 10)
	fmt.Printf("s1: %#v\ns2: %v\ns3: %v\n", s1, s2, s3)
}

//type Bar struct {
//}
//
//func convert(foos []Foo1) []Bar {
//	// 初始化bars
//	for _, foo := range foos {
//		if something() {
//			bars = append(bars, FooToBar(foo))
//
//		}
//	}
//	return bars
//}
