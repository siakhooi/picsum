package output

import (
	"errors"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestSaveImage_Success(t *testing.T) {
	// GIVEN
	tmpfile := "test_image.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	imageData := "fake image data"
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(imageData)),
	}

	// WHEN
	err := SaveImage(resp, tmpfile, false, false)
	if err != nil {
		t.Fatalf("SaveImage failed: %v", err)
	}

	// THEN
	data, err := os.ReadFile(tmpfile)
	if err != nil {
		t.Fatalf("Failed to read saved file: %v", err)
	}

	if string(data) != imageData {
		t.Errorf("Expected file content %q, got %q", imageData, string(data))
	}
}

func TestSaveImage_InvalidPath(t *testing.T) {
	// GIVEN
	// Use an invalid path (directory that doesn't exist)
	invalidPath := "/nonexistent/directory/image.jpg"

	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader("data")),
	}

	// WHEN
	err := SaveImage(resp, invalidPath, false, false)
	// THEN
	if err == nil {
		t.Error("Expected error for invalid path, got nil")
	}
}

func TestSaveImage_EmptyResponse(t *testing.T) {
	// GIVEN
	tmpfile := "test_empty.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader("")),
	}

	// WHEN
	err := SaveImage(resp, tmpfile, false, false)
	if err != nil {
		t.Fatalf("SaveImage failed: %v", err)
	}

	// THEN
	info, err := os.Stat(tmpfile)
	if err != nil {
		t.Fatalf("Failed to stat saved file: %v", err)
	}

	if info.Size() != 0 {
		t.Errorf("Expected empty file, got size %d", info.Size())
	}
}

// errorReader is a custom reader that always returns an error
type errorReader struct{}

func (e *errorReader) Read(_ []byte) (int, error) {
	return 0, errors.New("simulated read error")
}

func (e *errorReader) Close() error {
	return nil
}

func TestSaveImage_CopyError(t *testing.T) {
	// GIVEN
	tmpfile := "test_copy_error.jpg"
	defer func() { _ = os.Remove(tmpfile) }()

	resp := &http.Response{
		Body: &errorReader{},
	}

	// WHEN
	err := SaveImage(resp, tmpfile, false, false)
	if err == nil {
		t.Error("Expected error from io.Copy failure, got nil")
	}

	// THEN
	if !strings.Contains(err.Error(), "failed to save image") {
		t.Errorf("Expected error message to contain 'failed to save image', got: %v", err)
	}
}
