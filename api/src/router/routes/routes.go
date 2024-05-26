package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// All API routes
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func Settings(r *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
