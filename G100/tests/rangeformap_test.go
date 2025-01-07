package tests

import (
	"fmt"
	"testing"
)

func TestRangeForMap(t *testing.T) {
	m1 := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	m2 := copyMap(m1) // 创建初始map的拷贝

	for k, v := range m1 {
		m2[k] = v
		if v {
			m2[k+10] = true
		}
	}
	fmt.Println("m1:", m1)
	fmt.Println("m2:", m2)
}

func copyMap(src map[int]bool) map[int]bool {
	dst := make(map[int]bool)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
