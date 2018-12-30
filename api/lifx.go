package lifx

import "net/http"

type clientInstance struct {
	hostURL  string
	oauthKey string
	// Re-use the same client so TCP connections can be cached
	// http.Client is safe for re-use across goroutines
	client *http.Client
}

// Client is the interface required to interact with the LIFX HTTP API
type Client interface {
	ListLights()
	SetState()
	BreatheEffect()
}
