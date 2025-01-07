package errortest

import (
	"errors"
	"fmt"
	"net/http"
)

// 定义自定义错误类型 transientError
type transientError struct {
	err error
}

func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}

// 模拟从数据库中获取交易金额的函数
func getTransactionAmountFromDB(transactionID string) (float32, error) {
	// 这里模拟数据库错误，返回 transientError
	return 0, transientError{err: errors.New("database is down")}
}

/* 错误类型检查方式一 */

// getTransactionAmount1 使用方式一进行错误处理
func getTransactionAmount1(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}

	amount, err := getTransactionAmountFromDB(transactionID)
	if err != nil {
		return 0, transientError{err: err}
	}
	return amount, nil
}

// handler1 使用方式一的 HTTP 处理函数
func handler1(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transactionID")

	amount, err := getTransactionAmount1(transactionID)
	if err != nil {
		// 使用类型断言检查错误类型
		switch err := err.(type) {
		case transientError:
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// 返回正确响应
	fmt.Fprintf(w, "Transaction amount for ID %s is %.2f", transactionID, amount)
}

/* 错误类型检查方式二 */

// getTransactionAmount2 使用方式二进行错误处理
func getTransactionAmount2(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s", transactionID)
	}

	amount, err := getTransactionAmountFromDB(transactionID)
	if err != nil {
		// 使用 fmt.Errorf 包装原始错误，使用 %w 以便错误链形式传递
		return 0, fmt.Errorf("failed to get transaction %s: %w", transactionID, err)
	}
	return amount, nil
}

// handler2 使用方式二的 HTTP 处理函数
func handler2(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transactionID")

	amount, err := getTransactionAmount2(transactionID)
	if err != nil {
		var tErr transientError
		// 使用 errors.As() 检查错误链中是否包含 transientError 类型的错误
		if errors.As(err, &tErr) {
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	// 返回正确响应
	fmt.Fprintf(w, "Transaction amount for ID %s is %.2f", transactionID, amount)
}
