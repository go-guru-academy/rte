package rte

import (
	"net/http"
)

// Default is the default request input
type Default struct {
	W http.ResponseWriter
	R *http.Request
	// requestStart is timestamp for the request start
	// Unix epoch in nanoseconds
	requestStart int64
	// requestEnd is the timestamp for the request end
	// Unix epoch in nanoseconds
	requestEnd int64
	// requestDuration is the request duration
	requestDuration int64
}

// Start returns the time the request started
func (d *Default) Start() int64 {
	return d.requestStart
}

// End returns the time the request ended
func (d *Default) End() int64 {
	return d.requestEnd
}

// Duration returns the request duration
func (d *Default) Duration() int64 {
	return d.requestDuration
}
