// loc.go
package loc

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

// loc 表示一个位置服务。
type loc struct {
	httpClient *http.Client
	endpoint   string
}

// NewLoc 创建一个新的 loc 实例。
func NewLoc(endpoint string) *loc {
	return &loc{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		endpoint: endpoint,
	}
}

// validateAddress 检查提供的地址是否有效。
func (l *loc) validateAddress(address string) bool {
	// 简单的验证：检查地址是否为空。
	return address != ""
}

// getCoordinates 返回给定地址的纬度和经度。
func (l *loc) getCoordinates(ctx context.Context, address string) (lat, lng float32, err error) {
	// 验证地址。
	isValid := l.validateAddress(address)
	if !isValid {
		err = errors.New("无效的地址")
		return
	}

	// 检查上下文是否已被取消。
	if ctx.Err() != nil {
		err = ctx.Err()
		return
	}

	// 构建请求。
	req, err := http.NewRequestWithContext(ctx, "GET", l.endpoint, nil)
	if err != nil {
		return
	}

	// 添加查询参数。
	q := req.URL.Query()
	q.Add("address", address)
	req.URL.RawQuery = q.Encode()

	// 发送请求。
	resp, err := l.httpClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// 检查 HTTP 错误。
	if resp.StatusCode != http.StatusOK {
		err = errors.New("获取坐标失败")
		return
	}

	// 解析响应。
	var result struct {
		Lat float32 `json:"lat"`
		Lng float32 `json:"lng"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return
	}

	lat = result.Lat
	lng = result.Lng
	return
}
