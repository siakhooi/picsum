/*
Package download to download images from URLs
*/
package download

import (
	"fmt"
	"net/http"

	"github.com/siakhooi/picsum/internal/console"
)

/*
Image downloads an image from the given URL and returns the HTTP response
*/
func Image(url string, quiet bool) (*http.Response, error) {
	if !quiet {
		console.Stdoutln("Downloading from %s...", url)
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("server returned status: %s", resp.Status)
	}

	return resp, nil
}
