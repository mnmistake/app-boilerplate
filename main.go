package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/api"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    api.RootQuery,
	Mutation: api.RootMutation,
})

func main() {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})
	development := os.Getenv("ENV") == "development"

	server.InitDb()
	defer server.DB.Close()

	if development {
		http.Handle("/", http.FileServer(http.Dir("./client")))
	} // only serve static files in development via this server. Nginx is used in production instead

	http.Handle("/graphql", h) // todo: add auth middleware -> http.Handle("/graphql", server.RequireAuth(h))
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
