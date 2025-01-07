package tests

import (
	"errors"
	"fmt"
	"testing"
)

// 定义一个哨兵错误
var ErrNoRows = errors.New("no rows found")

// 模拟一个查询函数，根据id返回不同的错误
func query(id int) error {
	if id == 0 {
		// 直接返回哨兵错误
		return ErrNoRows
	} else if id == 1 {
		// 使用fmt.Errorf和%w包装哨兵错误
		return fmt.Errorf("query failed: %w", ErrNoRows)
	}
	// 无错误
	return nil
}

// 测试错误比较
func TestQueryErrorComparison(t *testing.T) {
	// 测试场景1：直接返回哨兵错误
	err := query(0)

	// 使用==操作符比较
	if err == ErrNoRows {
		t.Log("使用 == ：错误等于 ErrNoRows")
	} else {
		t.Error("使用 == ：错误不等于 ErrNoRows")
	}

	// 使用errors.Is比较
	if errors.Is(err, ErrNoRows) {
		t.Log("使用 errors.Is ：错误是 ErrNoRows")
	} else {
		t.Error("使用 errors.Is ：错误不是 ErrNoRows")
	}

	// 测试场景2：返回包装后的哨兵错误
	err = query(1)

	// 使用==操作符比较
	if err == ErrNoRows {
		t.Error("使用 == ：错误等于 ErrNoRows（不符合预期）")
	} else {
		t.Log("使用 == ：错误不等于 ErrNoRows（符合预期）")
	}

	// 使用errors.Is比较
	if errors.Is(err, ErrNoRows) {
		t.Log("使用 errors.Is ：错误是 ErrNoRows（符合预期）")
	} else {
		t.Error("使用 errors.Is ：错误不是 ErrNoRows")
	}

	// 测试场景3：无错误
	err = query(2)
	if err != nil {
		t.Error("预期没有错误，但得到：", err)
	} else {
		t.Log("没有错误返回（符合预期）")
	}
}
