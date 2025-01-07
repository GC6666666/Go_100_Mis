package tests

import (
	"math/rand"
	"runtime"
	"testing"
	"time"
)

var src = rand.NewSource(time.Now().UnixNano())

// randBytes 生成一个包含 128 个随机字节的数组
func randBytes() [129]byte {
	var b [129]byte
	r := rand.New(src)
	r.Read(b[:])
	return b
}

// TestMap1 测试 map 的内存使用情况
func TestMap1(t *testing.T) {
	m := make(map[int][129]byte)
	printMemUsage("before alloc")
	for i := 0; i < 1_000_000; i++ {
		m[i] = randBytes()
	}
	printMemUsage("after alloc")

	for i := 0; i < 1_000_000; i++ {
		delete(m, i)
	}
	printMemUsage("after delete")
	runtime.GC()
	runtime.KeepAlive(m)
	printMemUsage("after GC")
}
