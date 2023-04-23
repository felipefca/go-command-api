package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func Configuration(r *mux.Router) *mux.Router {
	routes := commandRoute

	for _, route := range routes {
		r.HandleFunc(route.Uri, route.Function).Methods(route.Method)
	}

	return r
}
