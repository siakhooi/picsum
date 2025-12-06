/*
Package output to save pic file
*/
package output

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
SaveImage saves the HTTP response body to a file with the given filename
*/
func SaveImage(resp *http.Response, filename string, quiet bool) error {
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
