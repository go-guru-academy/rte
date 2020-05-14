package rte

import (
	"net/http"

	"github.com/go-guru-academy/rte/routers/chi"
	"github.com/go-guru-academy/rte/routers/gorilla"
)

const (
	GORILLA int = iota + 1 // Start at 1 to error on no input
	CHI
)

type (
	Handler    func(*Default, interface{})
	Middleware func(func(*Default, interface{})) Handler
)

type Rte struct {
	router  Router
	handler http.Handler
}

type Router interface {
	AddRoute(method string, path string, handler http.HandlerFunc)
}

func New(routerPkg int) *Rte {
	r := &Rte{}
	switch routerPkg {
	case GORILLA:
		r.router, r.handler = gorilla.New()
	case CHI:
		r.router, r.handler = chi.New()
	default:
		panic("rte-error: invalid router type")
	}
	return r
}

// Returns the configured router instance
func (r *Rte) GetRouter() interface{} {
	return r.handler
}

func (r *Rte) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	r.handler.ServeHTTP(responseWriter, request)
}

func (r *Rte) addRoute(route *Route) {
	r.router.AddRoute(route.method, route.path, route.first(route.chainMiddleware(0)))
}

func (r *Rte) Post(path string, handler Handler, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodPost,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

func (r *Rte) Get(path string, handler Handler, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodGet,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

func (r *Rte) Put(path string, handler Handler, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodPut,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

func (r *Rte) Delete(path string, handler Handler, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodDelete,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

func (r *Rte) Connect(path string, handler Handler, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodConnect,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

func (r *Rte) Head(path string, handler Handler, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodHead,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

func (r *Rte) Options(path string, handler Handler, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodOptions,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

func (r *Rte) Patch(path string, handler Handler, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodPatch,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}

func (r *Rte) Trace(path string, handler Handler, input interface{}, middleware []Middleware) {
	r.addRoute(&Route{
		path:       path,
		method:     http.MethodTrace,
		middleware: middleware,
		input:      input,
		handler:    handler,
	})
}
