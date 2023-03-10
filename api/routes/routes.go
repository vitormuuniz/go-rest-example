package routes

import (
	"net/http"

	"go-rest-example/api/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	Path    string
	Method  string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, categoryRoutes CategoryRoutes, productRoutes ProductRoutes) {
	allRoutes := categoryRoutes.Routes()
	allRoutes = append(allRoutes, productRoutes.Routes()...)

	for _, route := range allRoutes {
		handler := middlewares.Logger(route.Handler)
		router.HandleFunc(route.Path, handler).Methods(route.Method)
	}
}
