/*
Package output to save pic file
*/
package output

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/*
promptForOverwrite asks the user for confirmation to overwrite a file.
Returns true if the user confirms, false otherwise.
*/
func promptForOverwrite(filename string, in io.Reader, out io.Writer) (bool, error) {
	if _, err := fmt.Fprintf(out, "File '%s' already exists. Overwrite? [y/N]: ", filename); err != nil {
		return false, fmt.Errorf("failed to write prompt: %v", err)
	}
	reader := bufio.NewReader(in)
	response, err := reader.ReadString('\n')
	if err != nil {
		return false, fmt.Errorf("failed to read user input: %v", err)
	}
	response = strings.TrimSpace(strings.ToLower(response))
	return response == "y" || response == "yes", nil
}

/*
SaveImage saves the HTTP response body to a file with the given filename
*/
func SaveImage(resp *http.Response, filename string, quiet bool, force bool) error {
	// Check if file exists
	if _, err := os.Stat(filename); err == nil {
		// File exists
		if !force {
			shouldOverwrite, err := promptForOverwrite(filename, os.Stdin, os.Stdout)
			if err != nil {
				return err
			}
			if !shouldOverwrite {
				return fmt.Errorf("user cancelled")
			}
		}
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer func() { _ = file.Close() }()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to save image: %v", err)
	}

	if !quiet {
		fmt.Printf("Image saved as %s\n", filename)
	}
	return nil
}
