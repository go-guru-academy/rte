package rte

import (
	"net/http"
)

// Default is the default request input
type Default struct {
	W http.ResponseWriter
	R *http.Request
	// start is timestamp for the request start
	// Unix epoch in nanoseconds
	start int64
	// end is the timestamp for the request end
	// Unix epoch in nanoseconds
	end int64
	// duration is the request duration
	duration int64
}

// Start returns the time the request started
func (d *Default) Start() int64 {
	return d.start
}

// End returns the time the request ended
func (d *Default) End() int64 {
	return d.end
}

// Duration returns the request duration
func (d *Default) Duration() int64 {
	return d.duration
}