package main

import (
	"net/http"
)

// Conn represents a client connection to a specific address
type Conn struct {
	address string
	// Re-use the same http.Client across requests so that TCP connections can be cached.
	// http.Client is safe for re-use across go-routines.
	client *http.Client
}
