package output

import (
	"bytes"
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

func TestPromptForOverwrite_UserConfirmsWithY(t *testing.T) {
	// GIVEN
	filename := "test.jpg"
	oldStdin := os.Stdin
	oldStdout := os.Stdout
	defer func() {
		os.Stdin = oldStdin
		os.Stdout = oldStdout
	}()

	r, w, _ := os.Pipe()
	os.Stdin = r
	rOut, wOut, _ := os.Pipe()
	os.Stdout = wOut

	_, _ = w.Write([]byte("y\n"))
	_ = w.Close()

	// WHEN
	result, err := promptForOverwrite(filename)

	_ = wOut.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, rOut)

	// THEN
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !result {
		t.Error("Expected true when user enters 'y'")
	}
	expectedPrompt := "File 'test.jpg' already exists. Overwrite? [y/N]: "
	if buf.String() != expectedPrompt {
		t.Errorf("Expected prompt %q, got %q", expectedPrompt, buf.String())
	}
}

func TestPromptForOverwrite_UserConfirmsWithYes(t *testing.T) {
	// GIVEN
	filename := "test.jpg"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	_, _ = w.Write([]byte("yes\n"))
	_ = w.Close()

	// WHEN
	result, err := promptForOverwrite(filename)

	// THEN
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !result {
		t.Error("Expected true when user enters 'yes'")
	}
}

func TestPromptForOverwrite_UserConfirmsWithUppercase(t *testing.T) {
	// GIVEN
	filename := "test.jpg"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	_, _ = w.Write([]byte("YES\n"))
	_ = w.Close()

	// WHEN
	result, err := promptForOverwrite(filename)

	// THEN
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !result {
		t.Error("Expected true when user enters 'YES'")
	}
}

func TestPromptForOverwrite_UserDeclinesWithN(t *testing.T) {
	// GIVEN
	filename := "test.jpg"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	_, _ = w.Write([]byte("n\n"))
	_ = w.Close()

	// WHEN
	result, err := promptForOverwrite(filename)

	// THEN
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result {
		t.Error("Expected false when user enters 'n'")
	}
}

func TestPromptForOverwrite_UserDeclinesWithNo(t *testing.T) {
	// GIVEN
	filename := "test.jpg"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	_, _ = w.Write([]byte("no\n"))
	_ = w.Close()

	// WHEN
	result, err := promptForOverwrite(filename)

	// THEN
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result {
		t.Error("Expected false when user enters 'no'")
	}
}

func TestPromptForOverwrite_UserDeclinesWithEmptyInput(t *testing.T) {
	// GIVEN
	filename := "test.jpg"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	_, _ = w.Write([]byte("\n"))
	_ = w.Close()

	// WHEN
	result, err := promptForOverwrite(filename)

	// THEN
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result {
		t.Error("Expected false (default No) when user enters empty input")
	}
}

func TestPromptForOverwrite_UserEntersInvalidInput(t *testing.T) {
	// GIVEN
	filename := "test.jpg"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	_, _ = w.Write([]byte("maybe\n"))
	_ = w.Close()

	// WHEN
	result, err := promptForOverwrite(filename)

	// THEN
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result {
		t.Error("Expected false when user enters invalid input")
	}
}

func TestPromptForOverwrite_UserEntersYesWithWhitespace(t *testing.T) {
	// GIVEN
	filename := "test.jpg"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	_, _ = w.Write([]byte("  yes  \n"))
	_ = w.Close()

	// WHEN
	result, err := promptForOverwrite(filename)

	// THEN
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !result {
		t.Error("Expected true when user enters 'yes' with whitespace")
	}
}

func TestPromptForOverwrite_ReadError(t *testing.T) {
	// GIVEN
	filename := "test.jpg"
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	// Mock stdin with an error reader
	r, w, _ := os.Pipe()
	os.Stdin = r
	_ = w.Close() // Close immediately to cause EOF

	// WHEN
	result, err := promptForOverwrite(filename)

	// THEN
	// Note: console.ReadLine() treats EOF as non-error, so it returns "" with nil error
	// The promptForOverwrite will return false (default No) for empty input
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result {
		t.Error("Expected false when EOF occurs (empty input)")
	}
}
