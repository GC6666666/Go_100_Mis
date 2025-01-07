package No_79

import (
	"io"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestHttp(t *testing.T) {
	var filename string
	// 打开文件
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Printf("failed to close file: %v\n", err)
		}
	}()
}

type handler struct {
	client http.Client
	url    string
}

func (h handler) getBody() (string, error) {
	resp, err := h.client.Get(h.url)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()

	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (h handler) getStatusCode(body io.Reader) (int, error) {
	resp, err := h.client.Post(h.url, "application/json", body)
	if err != nil {
		return 0, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("failed to close response: %v\n", err)
		}
	}()
	return resp.StatusCode, nil
}

func writeToFile(filename string, content []byte) (err error) {
	// 打开文件
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close() //忽略潜在错误
	}()
	_, err = f.Write(content)
	if err != nil {
		return err
	}
	return f.Sync() //将提交写入到磁盘
}

func handler(w http.ResponseWriter, req *http.Request) {
	err := foo(req)
	if err != nil {
		http.Error(w, "foo", http.StatusInternalServerError)
	}

	_, _ = w.Write([]byte("all good"))
	w.WriteHeader(http.StatusCreated)
}
