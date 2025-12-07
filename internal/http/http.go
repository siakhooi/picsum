/*
Package http provides an abstraction layer for HTTP operations
*/
package http

import (
	"net/http"
)

// Client interface defines HTTP operations
type Client interface {
	Get(url string) (*http.Response, error)
}

// DefaultClient implements Client using the standard http package
type DefaultClient struct{}

// Get performs an HTTP GET request using http.Get
func (c *DefaultClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

// NewDefaultClient creates a new DefaultClient instance
func NewDefaultClient() Client {
	return &DefaultClient{}
}
