package gorilla

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Gorilla struct {
	router *mux.Router
}

func New() (*Gorilla, *mux.Router) {
	m := mux.NewRouter()
	return &Gorilla{
		router: m,
	}, m
}

func (g *Gorilla) AddRoute(method string, path string, handler http.HandlerFunc) {
	g.router.HandleFunc(path, handler).Methods(method)
}
