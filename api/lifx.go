package lifx

import (
	"net/http"
)

// Client specifies the settings for communicating with the LIFX API
type Client struct {
	hostURL  string
	oauthKey string
	// Re-use the same client so TCP connections can be cached
	// http.Client is safe for re-use across goroutines
	client *http.Client
}
