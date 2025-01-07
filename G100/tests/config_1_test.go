package tests

import (
	"G100/5-Avoiding_interface_pollution/config"
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	// 没有限制作用
	// cfg := &config.IntConfig{Value: 42}
	var cfg config.IntConfigGetter = &config.IntConfig{Value: 42}
	value := cfg.Get()
	fmt.Println(value)
	// cfg.Set(43)
}
