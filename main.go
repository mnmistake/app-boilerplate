package main

import (
	"os"
	"net/http"
	
	"github.com/gorilla/handlers"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	//"github.com/labstack/echo"
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
	//e := echo.New()

	//e.Static("/", "./client")

	server.InitDb()
	defer server.DB.Close()

	http.Handle("/graphql", h)
	http.Handle("/", http.FileServer(http.Dir("./client")))
	http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	//e.Start(":8000")
}
