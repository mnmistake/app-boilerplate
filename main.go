package main

import (
	"net/http"

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
	http.ListenAndServe(":8000", nil)
	//e.Start(":8000")
}
