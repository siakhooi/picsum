package download

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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
	resp, err := Image(server.URL)

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
	resp, err := Image(server.URL)

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
	resp, err := Image(invalidURL)

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
	resp, err := Image(server.URL)

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
