package download

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	httpwrapper "github.com/siakhooi/picsum/internal/http"
)

func TestImage_Success(t *testing.T) {
	// GIVEN
	// Create a test server that returns 200 OK
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("fake image data"))
	}))
	defer server.Close()

	// WHEN
	resp, err := Image(server.URL, false)

	// THEN
	if err != nil {
		t.Fatalf("Image failed: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestImage_NonOKStatus(t *testing.T) {
	// GIVEN
	// Create a test server that returns 404 Not Found
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	// WHEN
	resp, err := Image(server.URL, false)

	// THEN
	if err == nil {
		t.Error("Expected error for non-OK status, got nil")
	}

	if resp != nil {
		t.Error("Expected nil response for non-OK status")
	}

	if !strings.Contains(err.Error(), "server returned status") {
		t.Errorf("Expected error message to contain 'server returned status', got: %v", err)
	}
}

func TestImage_InvalidURL(t *testing.T) {
	// GIVEN
	invalidURL := "http://invalid-host-that-does-not-exist-12345.com"

	// WHEN
	resp, err := Image(invalidURL, false)

	// THEN
	if err == nil {
		t.Error("Expected error for invalid URL, got nil")
	}

	if resp != nil {
		t.Error("Expected nil response for invalid URL")
	}

	if !strings.Contains(err.Error(), "failed to download image") {
		t.Errorf("Expected error message to contain 'failed to download image', got: %v", err)
	}
}

func TestImage_ServerError(t *testing.T) {
	// GIVEN
	// Create a test server that returns 500 Internal Server Error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	// WHEN
	resp, err := Image(server.URL, false)

	// THEN
	if err == nil {
		t.Error("Expected error for server error status, got nil")
	}

	if resp != nil {
		t.Error("Expected nil response for server error")
	}

	if !strings.Contains(err.Error(), "500") {
		t.Errorf("Expected error message to contain status code, got: %v", err)
	}
}

// MockHTTPClient is a mock implementation for testing
type MockHTTPClient struct {
	GetFunc func(url string) (*http.Response, error)
}

func (m *MockHTTPClient) Get(url string) (*http.Response, error) {
	if m.GetFunc != nil {
		return m.GetFunc(url)
	}
	return nil, nil
}

func TestImageWithClient_MockSuccess(t *testing.T) {
	// GIVEN
	mockClient := &MockHTTPClient{
		GetFunc: func(_ string) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(strings.NewReader("mocked image data")),
			}, nil
		},
	}

	// WHEN
	resp, err := ImageWithClient(mockClient, "http://example.com/image.jpg", true)

	// THEN
	if err != nil {
		t.Fatalf("ImageWithClient failed: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestImageWithClient_MockError(t *testing.T) {
	// GIVEN
	mockClient := &MockHTTPClient{
		GetFunc: func(_ string) (*http.Response, error) {
			return nil, http.ErrHandlerTimeout
		},
	}

	// WHEN
	resp, err := ImageWithClient(mockClient, "http://example.com/image.jpg", true)

	// THEN
	if err == nil {
		t.Error("Expected error, got nil")
	}

	if resp != nil {
		t.Error("Expected nil response on error")
	}

	if !strings.Contains(err.Error(), "failed to download image") {
		t.Errorf("Expected error message to contain 'failed to download image', got: %v", err)
	}
}

func TestImageWithClient_MockNonOKStatus(t *testing.T) {
	// GIVEN
	mockClient := &MockHTTPClient{
		GetFunc: func(_ string) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Status:     "404 Not Found",
				Body:       io.NopCloser(strings.NewReader("")),
			}, nil
		},
	}

	// WHEN
	resp, err := ImageWithClient(mockClient, "http://example.com/image.jpg", true)

	// THEN
	if err == nil {
		t.Error("Expected error for non-OK status, got nil")
	}

	if resp != nil {
		t.Error("Expected nil response for non-OK status")
	}

	if !strings.Contains(err.Error(), "server returned status") {
		t.Errorf("Expected error message to contain 'server returned status', got: %v", err)
	}
}

func TestImageWithClient_DefaultClient(t *testing.T) {
	// GIVEN
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("test data"))
	}))
	defer server.Close()

	client := httpwrapper.NewDefaultClient()

	// WHEN
	resp, err := ImageWithClient(client, server.URL, true)

	// THEN
	if err != nil {
		t.Fatalf("ImageWithClient failed: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
