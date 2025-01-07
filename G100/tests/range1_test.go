package tests

import (
	"fmt"
	"testing"
)

func TestRangeSharedArray(t *testing.T) {
	original := []int{1, 2, 3}
	fmt.Printf("Original slice:%v\n\n", original)

	for i, v := range original {
		fmt.Printf("Before modification: i = %d, v = %d\n", i, v)

		// 修改原始切片
		if i < 2 {
			original[i+1] *= 10

		}

		fmt.Printf("After modification: i = %d, v = %d, original[%d] = %d\n",
			i, v, i, original[i])
		fmt.Println("Current original slice:", original)
		fmt.Println()
	}

	fmt.Println("Final original slice:", original)
}
