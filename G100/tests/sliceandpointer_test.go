package tests

import (
	"runtime"
	"testing"
)

type Foo1 struct {
	v []int
}

func TestSliceAndPointer(t *testing.T) {
	foos := make([]Foo1, 1_000)
	PrintMemUsage("before GC")
	for i := 0; i < len(foos); i++ {
		foos[i] = Foo1{
			v: make([]int, 1024),
		}
	}
	PrintMemUsage("after For")
	two := keepFirstTwoElementsOnly(foos)
	runtime.GC()
	PrintMemUsage("after GC")
	runtime.KeepAlive(two)
}

func keepFirstTwoElementsOnly(foos []Foo1) []Foo1 {
	return foos[:2]
}
