package main

import (
	"fmt"
	"log"

	"github.com/graphql-go/graphql"
)

var (
	id          int
	content     string
	isCompleted bool
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "rootQuery",
	Fields: graphql.Fields{
		"todo": &graphql.Field{
			Type:        todoType,
			Description: "return a todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// Query via Todo ID and return a Todo's ID, content and is_completed status
				queryID, _ := params.Args["id"].(int)
				todo := QueryTodo(queryID)

				return todo, nil
			},
		},
		"todoList": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "return all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				data, _ := db.Query("SELECT id, content FROM todos")

				defer data.Close()
				for data.Next() {
					err := data.Scan(&id, &content)
					if err != nil {
						log.Fatal(err)
					}
					fmt.Println(id, content)

					return TodoList2{
						Todo{
							ID:      id,
							Content: content,
						},
					}, nil
				}
				return TodoList2{}, nil
			},
		},
	},
})
