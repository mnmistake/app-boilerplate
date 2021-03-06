package api

import (
	"github.com/graphql-go/graphql"

	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/api/sheets"
	"github.com/raunofreiberg/kyrene/server/authentication"
)

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// Authentication
		"registerUser": &graphql.Field{
			Type:        UserType,
			Description: "Register user",
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				username := params.Args["username"].(string)
				password := params.Args["password"].(string)
				token, err := authentication.RegisterUser(username, password)

				if err != nil {
					return nil, err
				}

				return token, nil
			},
		},
		"loginUser": &graphql.Field{
			Type:        UserType,
			Description: "Login user",
			Args: graphql.FieldConfigArgument{
				"username": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"password": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				username := params.Args["username"].(string)
				password := params.Args["password"].(string)
				token, err := authentication.LoginUser(username, password)

				if err != nil {
					return nil, err
				}

				return token, nil
			},
		},
		// Sheets & segments
		"createSheet": &graphql.Field{
			Type:        SheetType,
			Description: "Creates a sheet along with segments and attaches it to a user",
			Args: graphql.FieldConfigArgument{
				"name": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"segments": &graphql.ArgumentConfig{
					Type: graphql.NewList(InputObjectSegmentType),
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				jwt := params.Context.Value("jwt").(string)
				_, err := server.ValidateJWT(jwt)

				if err != nil {
					return nil, err
				}

				claims, err := server.ParseToken(jwt)

				name, _ := params.Args["name"].(string)
				userID := int(claims["id"].(float64))
				segments, _ := params.Args["segments"].([]interface{})
				sheet, err := sheets.InsertSheet(name, userID, segments)

				if err != nil {
					return nil, err
				}

				return sheet, nil
			},
		},
		"deleteSheet": &graphql.Field{
			Type:        SheetType,
			Description: "Deletes a sheet along with segments",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				jwt := params.Context.Value("jwt").(string)
				_, err := server.ValidateJWT(jwt)

				if err != nil {
					return nil, err
				}

				claims, _ := server.ParseToken(jwt)
				userID := int(claims["id"].(float64))
				sheetID := params.Args["id"].(int)
				sheet, error := sheets.DeleteSheet(sheetID, userID)

				if error != nil {
					return nil, error
				}

				return sheet, nil
			},
		},
		/* "createSegment": &graphql.Field{
			Type:        SegmentType,
			Description: "Create a segment and attaches it to a sheet",
			Args: graphql.FieldConfigArgument{
				"sheetId": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
				"label": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"content": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				sheetID, _ := params.Args["sheetId"].(int)
				label, _ := params.Args["label"].(string)
				content, _ := params.Args["content"].(string)

				segment, err := segments.InsertSegment(sheetID, label, content)

				if err != nil {
					return nil, err
				}

				return segment, nil
			},
		}, */
		/*
			"updateTodo": &graphql.Field{
				Type:        TodoType, // return type
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
				Type:        TodoType,
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
				Type:        TodoType,
				Description: "Delete all todos",
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					todos, err := DeleteTodos()

					if err != nil {
						return nil, err
					}

					return todos, nil
				},
			}, */
	},
})
