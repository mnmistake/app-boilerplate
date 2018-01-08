package main

import (
	"log"
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
		log.Fatal(err)
	}
}

func init() {
	// temporary test data - replace with DB
	/*todo1 := Todo{1, "todo 1", true}
	todo2 := Todo{2, "todo 2", false}
	todo3 := Todo{3, "todo 3", false}
	TodoList = append(TodoList, todo1, todo2, todo3)*/
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
