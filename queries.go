package main

import (
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
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
				queryID, _ := params.Args["id"].(int)

				for _, todo := range TodoList {
					if todo.ID == queryID {
						return todo, nil
					}
				}

				return Todo{}, nil
			},
		},
		"todoList": &graphql.Field{
			Type:        graphql.NewList(todoType),
			Description: "return all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return TodoList, nil
			},
		},
	},
})
