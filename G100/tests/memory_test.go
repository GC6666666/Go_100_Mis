package tests

import (
	"fmt"
	"runtime"
	"testing"
)

func TestMemory(t *testing.T) {
	m := make([]int, 10_000_000)
	for i := range m {
		m[i] = i
	}
	PrintAlloc()
	PrintMemUsage("first alloc")
	l := m[:5:5]
	//m = nil
	runtime.GC()
	PrintMemUsage("second alloc")
	PrintAlloc()
	runtime.KeepAlive(l)
}
func PrintAlloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d KB\n", m.Alloc/1024)
}
