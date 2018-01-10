package main

import (
	"os"
	"net/http"
	
	"github.com/gorilla/handlers"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/raunofreiberg/kyrene/server"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    server.RootQuery,
	Mutation: server.RootMutation,
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

	http.Handle("/graphql", h)
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
