/*
Package download to download images from URLs
*/
package download

import (
	"fmt"
	"net/http"

	"github.com/siakhooi/picsum/internal/console"
	httpwrapper "github.com/siakhooi/picsum/internal/http"
)

/*
ImageWithClient downloads an image from the given URL using the provided HTTP client
*/
func ImageWithClient(client httpwrapper.Client, url string, quiet bool) (*http.Response, error) {
	if !quiet {
		console.Stdoutln("Downloading from %s...", url)
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to download image: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, fmt.Errorf("server returned status: %s", resp.Status)
	}

	return resp, nil
}

/*
Image downloads an image from the given URL and returns the HTTP response
Uses the default HTTP client
*/
func Image(url string, quiet bool) (*http.Response, error) {
	return ImageWithClient(httpwrapper.NewDefaultClient(), url, quiet)
}
