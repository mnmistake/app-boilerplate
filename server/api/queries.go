package api

import (
	"github.com/graphql-go/graphql"
	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/authentication"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"todo": &graphql.Field{
			Type:        TodoType,
			Description: "return a todo",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// Query via Todo ID and return a Todo's ID, content and is_completed status
				queryID, _ := params.Args["id"].(int)
				queriedTodo, err := QueryTodo(queryID)

				if err != nil {
					return nil, err
				}

				return queriedTodo, nil
			},
		},
		"todoList": &graphql.Field{
			Type:        graphql.NewList(TodoType),
			Description: "return all todos",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				res, err := server.RequireAuth(
					params.Context.Value("jwt").(string),
					QueryTodos,
				)

				if err != nil {
					return nil, err
				}

				return res, err
			},
		},
		"users": &graphql.Field{
			Type:        graphql.NewList(UserType),
			Description: "Return all users",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				res, err := server.RequireAuth(
					params.Context.Value("jwt").(string),
					authentication.QueryUsers,
				)

				if err != nil {
					return nil, err
				}

				return res, nil
			},
		},
	},
})
