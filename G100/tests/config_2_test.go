package tests

import (
	"G100/5-Avoiding_interface_pollution/config"
	"fmt"
	"testing"
)

type Foo struct {
	threshold config.IntConfigGetter
}

func CreateNewFoo(threshold config.IntConfigGetter) Foo {
	return Foo{threshold: threshold}
}

func (f Foo) Bar() int {
	threshold := f.threshold.Get()
	return threshold
}

func TestCreateNewFoo(t *testing.T) {
	cfg := &config.IntConfig{Value: 44}

	// 创建一个新的 Foo1 实例
	foo := CreateNewFoo(cfg)

	// 验证 Foo1 实例的 threshold 字段是否为传入的 IntConfigGetter
	if foo.threshold != cfg {
		t.Errorf("Expected foo.threshold to be %v, but got %v", cfg, foo.threshold)
	}

	// 调用 Foo1 实例的 Bar 方法,并验证返回值是否正确
	expectedThreshold := 44
	actualThreshold := foo.Bar()
	if actualThreshold != expectedThreshold {
		t.Errorf("Expected threshold to be %d, but got %d", expectedThreshold, actualThreshold)
	} else {
		fmt.Println("set success!!!")
	}

	// 修改 IntConfig 的值,并验证 Foo1 实例的 Bar 方法是否返回新的值
	cfg.Set(55)
	expectedThreshold = 55
	actualThreshold = foo.Bar()
	if actualThreshold != expectedThreshold {
		t.Errorf("Expected threshold to be %d, but got %d", expectedThreshold, actualThreshold)
	} else {
		fmt.Println("set success!!!")
	}
}
