package api

import (
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
				insertedTodo := InsertTodo(content)

				return insertedTodo, nil
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
				id := params.Args["id"].(int)
				modifiedTodo := UpdateTodo(id, IsCompleted)

				return modifiedTodo, nil
			},
		},
		"deleteTodo": &graphql.Field{
			Type:        todoType,
			Description: "Delete todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				id := params.Args["id"].(int)
				return DeleteTodo(id), nil
			},
		},
		"deleteTodos": &graphql.Field{
			Type:        todoType,
			Description: "Delete all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return DeleteTodos(), nil
			},
		},
	},
})
