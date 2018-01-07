package main

import (
	"math/rand"

	"github.com/graphql-go/graphql"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createTodo": &graphql.Field{
			Type:        todoType,
			Description: "create todo",
			Args: graphql.FieldConfigArgument{ // args that the mutation takes
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) { // handler of the mutation
				// handle creation of the todo
				content, _ := params.Args["content"].(string)
				newID := rand.Intn(100)
				newTodo := Todo{
					ID:          newID,
					Content:     content,
					IsCompleted: false,
				}
				TodoList = append(TodoList, newTodo)

				return newTodo, nil
			},
		},
		"updateTodo": &graphql.Field{
			Type:        todoType, // return type
			Description: "Update todo",
			Args: graphql.FieldConfigArgument{
				"isCompleted": &graphql.ArgumentConfig{
					Type: graphql.Boolean,
				},
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				IsCompleted, _ := params.Args["isCompleted"].(bool)
				id, _ := params.Args["id"].(int)
				targetTodo := Todo{}

				for _, todo := range TodoList {
					if todo.ID == id {
						todo.IsCompleted = IsCompleted
						targetTodo = todo
						break
					}
				}
				return targetTodo, nil
			},
		},
	},
})
