package tests

import (
	"fmt"
	"runtime"
	"testing"
)

// 模拟接收消息
func receiveMessage() []byte {
	return make([]byte, 1_000_000)
}

// 获取消息类型
func getMessageType(msg []byte) []byte {
	return msg[:5:5]
	//return append([]byte(nil), msg[:5]...)
}

// 存储消息类型
var messageTypes [][]byte

func storeMessageType(msgType []byte) {
	messageTypes = append(messageTypes, msgType)
	if len(messageTypes) > 1000 {
		messageTypes = messageTypes[1:]
	}
}

func consumeMessages(count int) {
	for i := 0; i < count; i++ {
		msg := receiveMessage()
		storeMessageType(getMessageType(msg))
		//msg = nil
	}
}

func TestMemoryLeak(t *testing.T) {
	// 初始内存使用
	PrintMemUsage("Initial memory usage")

	// 处理1000条消息
	consumeMessages(1000)

	messageTypes = nil
	// 强制GC
	runtime.GC()

	// 最终内存使用
	PrintMemUsage("Final memory usage")

	// 这里的问题是当我后续不使用messageTypes，仍旧没有被内存回收，内存的占用量仍旧很高

	// 验证
	//if len(messageTypes) != 1000 {
	//	t.Errorf("Expected 1000 message types, got %d", len(messageTypes))
	//}
	//
	//totalCapacity := 0
	//for _, msgType := range messageTypes {
	//	totalCapacity += cap(msgType)
	//}

	//fmt.Printf("Total capacity of stored message types: %d bytes\n", totalCapacity)
	//
	//if totalCapacity < 900_000_000 { // 预期接近1GB
	//	t.Errorf("Expected capacity close to 1GB, got %d bytes", totalCapacity)
	//}
}

func PrintMemUsage(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%s: Alloc = %v MiB, TotalAlloc = %v MiB, Sys = %v MiB\n",
		msg, BToMb(m.Alloc), BToMb(m.TotalAlloc), BToMb(m.Sys))
}

func BToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
