package console

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestStdout(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Stdout("Hello %s", "World")

	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	got := buf.String()
	want := "Hello World"
	if got != want {
		t.Errorf("Stdout() = %q, want %q", got, want)
	}
}

func TestStdoutln(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	Stdoutln("Hello %s", "World")

	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	got := buf.String()
	want := "Hello World\n"
	if got != want {
		t.Errorf("Stdoutln() = %q, want %q", got, want)
	}
}

func TestStderr(t *testing.T) {
	// Capture stderr
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	Stderr("Error %d", 404)

	_ = w.Close()
	os.Stderr = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	got := buf.String()
	want := "Error 404"
	if got != want {
		t.Errorf("Stderr() = %q, want %q", got, want)
	}
}

func TestStderrln(t *testing.T) {
	// Capture stderr
	old := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	Stderrln("Error %d", 404)

	_ = w.Close()
	os.Stderr = old

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	got := buf.String()
	want := "Error 404\n"
	if got != want {
		t.Errorf("Stderrln() = %q, want %q", got, want)
	}
}

func TestReadLine(t *testing.T) {
	// Mock stdin
	old := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	input := "test input\n"
	_, _ = w.Write([]byte(input))
	_ = w.Close()

	got, err := ReadLine()
	os.Stdin = old

	if err != nil {
		t.Fatalf("ReadLine() error = %v", err)
	}
	if got != input {
		t.Errorf("ReadLine() = %q, want %q", got, input)
	}
}

func TestReadAll(t *testing.T) {
	// Mock stdin
	old := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	input := "line 1\nline 2\nline 3"
	_, _ = w.Write([]byte(input))
	_ = w.Close()

	got, err := ReadAll()
	os.Stdin = old

	if err != nil {
		t.Fatalf("ReadAll() error = %v", err)
	}
	if got != input {
		t.Errorf("ReadAll() = %q, want %q", got, input)
	}
}

func TestScanner(t *testing.T) {
	// Mock stdin
	old := os.Stdin
	r, w, _ := os.Pipe()
	os.Stdin = r

	input := "line 1\nline 2\nline 3\n"
	_, _ = w.Write([]byte(input))
	_ = w.Close()

	scanner := Scanner()
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	os.Stdin = old

	got := strings.Join(lines, "\n")
	want := "line 1\nline 2\nline 3"
	if got != want {
		t.Errorf("Scanner() lines = %q, want %q", got, want)
	}
}
