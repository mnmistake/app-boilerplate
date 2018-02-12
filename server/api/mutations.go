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
				insertedTodo, err := InsertTodo(content)

				if err != nil {
					return nil, err
				}

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
				modifiedTodo, err := UpdateTodo(id, IsCompleted)

				if err != nil {
					return nil, err
				}

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
				todo, err := DeleteTodo(id)
				
				if err != nil {
					return nil, err
				}

				return todo, nil
			},
		},
		"deleteTodos": &graphql.Field{
			Type:        todoType,
			Description: "Delete all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				todos, err := DeleteTodos()

				if err != nil {
					return nil, err
				}
				
				return todos, nil
			},
		},
	},
})
