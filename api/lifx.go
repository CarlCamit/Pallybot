package lifx

import (
	"net"
	"net/http"
	"time"
)

// Client specifies the settings for communicating with the LIFX API
type Client struct {
	hostURL  string
	oauthKey string
	// Re-use the same client so TCP connections can be cached
	// http.Client is safe for re-use across goroutines
	client *http.Client
}

// NewClient returns a Client instance
func NewClient(hostURL string, oauthKey string) *Client {
	transport := &http.Transport{
		// http.DefaultTransport values
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,

		// Custom values
		MaxIdleConns:        1,
		MaxIdleConnsPerHost: 1,
	}

	return &Client{
		hostURL:  hostURL,
		oauthKey: oauthKey,
		client:   &http.Client{Transport: transport},
	}
}
