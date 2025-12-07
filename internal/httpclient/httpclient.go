/*
Package httpclient provides an abstraction layer for HTTP operations
*/
package httpclient

import (
	"net/http"
)

// Getter interface defines HTTP operations
type Getter interface {
	Get(url string) (*http.Response, error)
}

// DefaultClient implements Getter using the standard http package
type DefaultClient struct{}

// Get performs an HTTP GET request using http.Get
func (c *DefaultClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

// NewDefaultClient creates a new DefaultClient instance
func NewDefaultClient() Getter {
	return &DefaultClient{}
}
