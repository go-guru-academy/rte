package rte

import (
	"fmt"
	"net/http"
	"time"
)

type Route struct {
	// The request path
	path string

	// The request method
	method string

	// middleware can be used to modify requests and/or responses
	middleware []Middleware

	// input is custom input needed to handle the request
	// input is typically a struct pointer
	input interface{}

	// The request handler
	handler Handler
}

// first is a default middleware that populates the Default struct
// with values required to process requests
func (route *Route) first(next func(*Default, interface{})) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d := &Default{
			R:     r,
			W:     w,
			start: time.Now().UnixNano(),
		}
		fmt.Println(d)
		fmt.Println(route)
		next(d, route.input)
		d.end = time.Now().UnixNano()
		d.duration = d.end - d.start
	}
}

// chainMiddleware appends a middleware function to the chain.
func (route *Route) chainMiddleware(i int) Handler {
	if i == len(route.middleware) {
		return route.handler
	}
	f := route.middleware[i]
	return f(route.chainMiddleware(i + 1))
}
