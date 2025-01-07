package tests

import (
	"errors"
	"fmt"
	"testing"
)

func TestProcessRequest(t *testing.T) {
	// 模拟测试过程
	if err := processRequest(); err != nil {
		t.Errorf("ProcessRequest failed: %v", err)
	}
}

const (
	StatusSuccess          = "Success"
	StatusAuthFailed       = "AuthFailed"
	StatusProcessingFailed = "ProcessingFailed"
)

func processRequest() error {
	var status string

	// 使用匿名函数包装，以捕获最新的 status 值
	defer func() { notify(status) }()
	defer func() { incrementCounter(status) }()
	//defer notify(&status)
	//defer incrementCounter(&status)

	// 身份验证步骤
	if err := authenticate(); err != nil {
		status = StatusAuthFailed
		return err
	}

	// 数据处理步骤
	if err := processData(); err != nil {
		status = StatusProcessingFailed
		return err
	}

	// 所有步骤成功，设置状态为成功
	status = StatusSuccess
	return nil
}

func authenticate() error {
	// 实现身份验证逻辑
	// 这里模拟失败情况
	return errors.New("authentication failed")
}

func processData() error {
	// 实现数据处理逻辑
	// 这里假设数据处理成功
	return nil
}

func notify(status string) {
	// 根据状态发送通知
	fmt.Printf("Notification sent with status: %s\n", status)
}

func incrementCounter(status string) {
	// 根据状态更新计数器
	fmt.Printf("Counter incremented for status: %s\n", status)
}
