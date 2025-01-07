package tests

import (
	"fmt"
	"runtime"
	"testing"
)

const (
	megabyte     = 1024 * 1024
	sliceSize    = 10 * megabyte // 10MB
	iterations   = 100
	expectedLeak = sliceSize * iterations / megabyte // 预期泄漏的MB数
)

func TestMemoryLeak1(t *testing.T) {
	var leakedSlices [][]byte

	printMemUsage("Initial memory usage")

	for i := 0; i < iterations; i++ {
		bigSlice := make([]byte, sliceSize)
		leakedSlices = append(leakedSlices, bigSlice[:1]) // 只保留每个大切片的第一个字节
	}

	runtime.GC()
	printMemUsage("After creating slices and GC")

	leakSize := calculateLeakSize(leakedSlices)
	fmt.Printf("Calculated leak size: %d MB\n", leakSize/megabyte)

	if leakSize < expectedLeak*megabyte {
		t.Errorf("Expected memory leak of at least %d MB, but got %d MB", expectedLeak, leakSize/megabyte)
	}

	runtime.KeepAlive(leakedSlices) // 确保 leakedSlices 不会被过早回收
}

func printMemUsage(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: Alloc = %v MB, TotalAlloc = %v MB, Sys = %v MB\n",
		msg, bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys))
}

func bToMb(b uint64) uint64 {
	return b / megabyte
}

func calculateLeakSize(slices [][]byte) int {
	total := 0
	for _, s := range slices {
		total += cap(s)
	}
	return total
}
