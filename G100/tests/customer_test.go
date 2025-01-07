package tests

import (
	"G100/5-Avoiding_interface_pollution/customer"
	"encoding/json"
	"fmt"
	"testing"
)

// mockStore 是一个模拟的存储实现，用于测试
type mockStore struct {
	Customer []customer.Customer
	Err      error
}

func (ms *mockStore) StoreCustomer(c customer.Customer) error {
	if ms.Err != nil {
		return ms.Err
	}
	ms.Customer = append(ms.Customer, c)
	return nil
}
func TestCreateNewCustomer(t *testing.T) {
	// 创建一个模拟的存储实现
	mock := &mockStore{}

	// 创建 CustomerService 并注入模拟存储实现
	customerService := customer.CustomerService{Storer: mock}

	// 调用 CreateNewCustomer 方法
	for i := 0; i < 5; i++ {
		err := customerService.CreateNewCustomer(i)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
	}
	// 验证模拟存储实现是否正确存储了客户
	if len(mock.Customer) != 5 {
		t.Fatalf("expected 5 customer, got %d", len(mock.Customer))
	}

	fmt.Println(mock)

	jsonData, _ := json.Marshal(mock)
	fmt.Println(string(jsonData))
}
