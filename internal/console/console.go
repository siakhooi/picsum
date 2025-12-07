// Package console provides simple functions for interacting with standard input, output, and error streams.
package console

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Stdout writes to standard output
func Stdout(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, format, args...)
}

// Stdoutln writes to standard output with a newline
func Stdoutln(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stdout, format+"\n", args...)
}

// Stderr writes to standard error
func Stderr(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format, args...)
}

// Stderrln writes to standard error with a newline
func Stderrln(format string, args ...interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, format+"\n", args...)
}

// ReadLine reads a single line from standard input
func ReadLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}
	return line, nil
}

// ReadAll reads all content from standard input
func ReadAll() (string, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Scanner returns a scanner for reading from standard input line by line
func Scanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}
