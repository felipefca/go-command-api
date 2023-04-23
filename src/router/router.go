package router

import (
	"api/src/router/route"

	"github.com/gorilla/mux"

	httpSwagger "github.com/swaggo/http-swagger"

	_ "api/docs"
)

// Generate will return router with all endpoint configurations
func Generate() *mux.Router {
	r := mux.NewRouter()

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return route.Configuration(r)
}
