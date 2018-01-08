package main

import (
	"net/http"

	"github.com/graphql-go/handler"

	"github.com/graphql-go/graphql"
)

type Todo struct {
	ID          int    `json:"id,omitempty"`
	Content     string `json:"content,omitempty"`
	IsCompleted bool   `json:"isCompleted,omitempty"`
}

var TodoList []Todo

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	// temporary test data - replace with DB
}

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    RootQuery,
	Mutation: RootMutation,
})

func main() {
	h := handler.New(&handler.Config{
		Schema:   &Schema,
		Pretty:   true,
		GraphiQL: true,
	})

	InitDb()
	defer db.Close()
	http.Handle("/graphql", h)
	http.ListenAndServe(":8000", nil)
}
