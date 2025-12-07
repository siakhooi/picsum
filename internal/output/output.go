/*
Package output to save pic file
*/
package output

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/siakhooi/picsum/internal/console"
)

/*
promptForOverwrite asks the user for confirmation to overwrite a file.
Returns true if the user confirms, false otherwise.
*/
func promptForOverwrite(filename string) (bool, error) {
	console.Stdout("File '%s' already exists. Overwrite? [y/N]: ", filename)
	response, err := console.ReadLine()
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
			shouldOverwrite, err := promptForOverwrite(filename)
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
		console.Stdoutln("Image saved as %s", filename)
	}
	return nil
}
