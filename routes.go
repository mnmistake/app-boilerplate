package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Todos",
		"GET",
		"/todos",
		GetTodos,
	},
	Route{
		"CREATE single TODO",
		"POST",
		"/todos",
		PostTodo,
	},
}

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.Handler)
	}

	return router
}
