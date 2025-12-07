package http

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDefaultClient_Get_Success(t *testing.T) {
	// Create a test server that returns a successful response
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("test response"))
	}))
	defer server.Close()

	client := NewDefaultClient()
	resp, err := client.Get(server.URL)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	_ = resp.Body.Close()

	if string(body) != "test response" {
		t.Errorf("Expected body 'test response', got '%s'", string(body))
	}
}

func TestDefaultClient_Get_InvalidURL(t *testing.T) {
	client := NewDefaultClient()
	_, err := client.Get("http://invalid-domain-that-does-not-exist-12345.com")

	if err == nil {
		t.Error("Expected error for invalid URL, got nil")
	}
}

func TestDefaultClient_Get_StatusCodes(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
	}{
		{"OK", http.StatusOK},
		{"Not Found", http.StatusNotFound},
		{"Internal Server Error", http.StatusInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(tt.statusCode)
			}))
			defer server.Close()

			client := NewDefaultClient()
			resp, err := client.Get(server.URL)

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			defer func() { _ = resp.Body.Close() }()

			if resp.StatusCode != tt.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.statusCode, resp.StatusCode)
			}
		})
	}
}

func TestNewDefaultClient(t *testing.T) {
	client := NewDefaultClient()

	if client == nil {
		t.Fatal("NewDefaultClient() returned nil")
	}

	if _, ok := client.(*DefaultClient); !ok {
		t.Error("NewDefaultClient() did not return *DefaultClient")
	}
}

// MockClient is a mock implementation for testing
type MockClient struct {
	GetFunc func(url string) (*http.Response, error)
}

func (m *MockClient) Get(url string) (*http.Response, error) {
	if m.GetFunc != nil {
		return m.GetFunc(url)
	}
	return nil, nil
}

func TestMockClient_Usage(t *testing.T) {
	// Example of how to use MockClient in tests
	mock := &MockClient{
		GetFunc: func(_ string) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("mocked response")),
			}, nil
		},
	}

	resp, err := mock.Get("http://example.com")

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)
	if string(body) != "mocked response" {
		t.Errorf("Expected 'mocked response', got '%s'", string(body))
	}
}
