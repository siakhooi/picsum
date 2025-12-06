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
SaveImage saves the HTTP response body to a file with the given filename
*/
func SaveImage(resp *http.Response, filename string, quiet bool, force bool) error {
	// Check if file exists
	if _, err := os.Stat(filename); err == nil {
		// File exists
		if !force {
			// Prompt for confirmation (default: No)
			fmt.Printf("File '%s' already exists. Overwrite? [y/N]: ", filename)
			reader := bufio.NewReader(os.Stdin)
			response, err := reader.ReadString('\n')
			if err != nil {
				return fmt.Errorf("failed to read user input: %v", err)
			}
			response = strings.TrimSpace(strings.ToLower(response))
			if response != "y" && response != "yes" {
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
