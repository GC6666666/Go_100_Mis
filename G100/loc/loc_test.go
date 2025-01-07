// loc_test.go
package loc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

// mockServer 模拟一个地理编码 API，用于测试。
func mockServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		address := r.URL.Query().Get("address")

		switch address {
		case "1600 Amphitheatre Parkway, Mountain View, CA":
			// 模拟成功响应。
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]float32{
				"lat": 37.4220,
				"lng": -122.0841,
			})
		case "":
			// 模拟无效地址响应。
			http.Error(w, "Bad Request: Invalid address", http.StatusBadRequest)
		default:
			// 模拟地址未找到响应。
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	})

	return httptest.NewServer(handler)
}

func TestGetCoordinates(t *testing.T) {
	// 设置模拟服务器。
	server := mockServer()
	defer server.Close()

	// 创建一个新的 loc 实例，使用模拟服务器的 URL。
	l := NewLoc(server.URL)

	// 测试用例。
	tests := []struct {
		name      string
		address   string
		expectLat float32
		expectLng float32
		expectErr bool
	}{
		{
			name:      "有效地址",
			address:   "1600 Amphitheatre Parkway, Mountain View, CA",
			expectLat: 37.4220,
			expectLng: -122.0841,
			expectErr: false,
		},
		{
			name:      "空地址",
			address:   "",
			expectLat: 0,
			expectLng: 0,
			expectErr: true,
		},
		{
			name:      "地址未找到",
			address:   "未知地址",
			expectLat: 0,
			expectLng: 0,
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			lat, lng, err := l.getCoordinates(ctx, tt.address)

			if tt.expectErr {
				if err == nil {
					t.Errorf("测试 %q：预期错误，但得到 nil", tt.name)
				}
			} else {
				if err != nil {
					t.Errorf("测试 %q：出现意外错误：%v", tt.name, err)
					return
				}
				if lat != tt.expectLat || lng != tt.expectLng {
					t.Errorf("测试 %q：坐标不匹配：得到 (%v, %v)，期望 (%v, %v)", tt.name, lat, lng, tt.expectLat, tt.expectLng)
				}
			}
		})
	}
}

func TestGetCoordinates_ContextCanceled(t *testing.T) {
	// 设置模拟服务器。
	server := mockServer()
	defer server.Close()

	// 创建一个新的 loc 实例。
	l := NewLoc(server.URL)

	// 创建一个立即被取消的上下文。
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // 立即取消上下文。

	lat, lng, err := l.getCoordinates(ctx, "1600 Amphitheatre Parkway, Mountain View, CA")

	if err == nil {
		t.Errorf("预期因上下文取消而出错，但得到 nil")
	} else if !errors.Is(err, context.Canceled) {
		t.Errorf("预期 context.Canceled 错误，得到：%v", err)
	}

	if lat != 0 || lng != 0 {
		t.Errorf("由于错误，预期坐标为 (0, 0)，但得到 (%v, %v)", lat, lng)
	}
}
