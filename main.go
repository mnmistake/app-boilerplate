package main

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"

	"github.com/labstack/echo"
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
	e := echo.New()

	e.Static("/", "../client")
	e.File("/bundle.js", "../dist")

	server.InitDb()
	http.Handle("/graphql", h)
	e.Start(":8000")
}
