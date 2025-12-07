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

	// Mock stdin with an error reader that returns a non-EOF error
	// Close the read end to cause a read error
	r, w, _ := os.Pipe()
	_ = r.Close() // Close read end to cause error on read
	os.Stdin = r
	_, _ = w.Write([]byte("y\n"))
	_ = w.Close()

	// WHEN
	result, err := promptForOverwrite(filename)

	// THEN
	// Should get an error because reading from closed pipe returns error
	if err == nil {
		t.Error("Expected error when reading from closed pipe")
	}
	if result {
		t.Error("Expected false when error occurs")
	}
	if err != nil && !strings.Contains(err.Error(), "failed to read user input") {
		t.Errorf("Expected error message to contain 'failed to read user input', got: %v", err)
	}
}

func TestSaveImage_FileExistsAndUserDeclines(t *testing.T) {
	// GIVEN
	tmpfile := "test_decline.jpg"
	// Create an existing file
	err := os.WriteFile(tmpfile, []byte("existing content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer func() { _ = os.Remove(tmpfile) }()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r
	_, _ = w.Write([]byte("n\n")) // User declines
	_ = w.Close()

	imageData := "new image data"
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(imageData)),
	}

	// WHEN
	err = SaveImage(resp, tmpfile, true, false)

	// THEN
	if err == nil {
		t.Error("Expected error when user declines overwrite")
	}
	if err != nil && !strings.Contains(err.Error(), "user cancelled") {
		t.Errorf("Expected error message to contain 'user cancelled', got: %v", err)
	}

	// Verify original file is unchanged
	data, _ := os.ReadFile(tmpfile)
	if string(data) != "existing content" {
		t.Error("Original file should not be modified when user declines")
	}
}

func TestSaveImage_FileExistsAndPromptError(t *testing.T) {
	// GIVEN
	tmpfile := "test_prompt_error.jpg"
	// Create an existing file
	err := os.WriteFile(tmpfile, []byte("existing content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer func() { _ = os.Remove(tmpfile) }()

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	// Create a closed pipe to trigger read error
	r, w, _ := os.Pipe()
	_ = r.Close() // Close read end to cause error
	os.Stdin = r
	_, _ = w.Write([]byte("y\n"))
	_ = w.Close()

	imageData := "new image data"
	resp := &http.Response{
		Body: io.NopCloser(strings.NewReader(imageData)),
	}

	// WHEN
	err = SaveImage(resp, tmpfile, true, false)

	// THEN
	if err == nil {
		t.Error("Expected error when prompt fails")
	}
	if err != nil && !strings.Contains(err.Error(), "failed to read user input") {
		t.Errorf("Expected error message to contain 'failed to read user input', got: %v", err)
	}
}
