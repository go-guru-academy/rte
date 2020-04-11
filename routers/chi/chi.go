package chi

import (
	"net/http"

	chipkg "github.com/go-chi/chi"
)

type Chi struct {
	router *chipkg.Mux
}

func New() (*Chi, *chipkg.Mux) {
	m := chipkg.NewRouter()
	return &Chi{
		router: m,
	}, m
}

func (c *Chi) AddRoute(method string, path string, handler http.HandlerFunc) {
	c.router.HandleFunc(path, handler)
}
