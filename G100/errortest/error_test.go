package errortest

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler1_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/?transactionID=12345", nil)
	w := httptest.NewRecorder()

	handler1(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("Expected status %d, got %d", http.StatusServiceUnavailable, resp.StatusCode)
	} else {
		t.Errorf("Expected status %d, got %d", http.StatusServiceUnavailable, resp.StatusCode)
	}
	body := w.Body.String()
	if !strings.Contains(body, "transient error") {
		t.Errorf("Expected error message containing 'transient error', got '%s'", body)
	}
}

func TestHandler2_Success(t *testing.T) {
	req := httptest.NewRequest("GET", "/?transactionID=12345", nil)
	w := httptest.NewRecorder()

	handler2(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Errorf("Expected status %d, got %d", http.StatusServiceUnavailable, resp.StatusCode)
	}
	body := w.Body.String()
	if !strings.Contains(body, "failed to get transaction") {
		t.Errorf("Expected error message containing 'failed to get transaction', got '%s'", body)
	}
}

func TestHandler1_InvalidID(t *testing.T) {
	req := httptest.NewRequest("GET", "/?transactionID=12", nil)
	w := httptest.NewRecorder()

	handler1(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}

func TestHandler2_InvalidID(t *testing.T) {
	req := httptest.NewRequest("GET", "/?transactionID=12", nil)
	w := httptest.NewRecorder()

	handler2(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
	}
}
