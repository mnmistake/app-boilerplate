package api

import (
	"github.com/graphql-go/graphql"

	"github.com/raunofreiberg/kyrene/server/api/segments"
	"github.com/raunofreiberg/kyrene/server/api/users"
	"github.com/raunofreiberg/kyrene/server/model"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"username": &graphql.Field{
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var SheetType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Sheet",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"userId": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
		"segments": &graphql.Field{
			Type:        graphql.NewList(SegmentType),
			Description: "Query all segments related to a sheet",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				/* res, err := server.RequireAuth(
					params.Context.Value("jwt").(string),
					QueryTodos,
				) */

				sheetID := params.Source.(model.Sheet).ID
				res, err := segments.QuerySegments(sheetID)

				if err != nil {
					return nil, err
				}

				return res, err
			},
		},
		"user": &graphql.Field{
			Type:        UserType,
			Description: "Query the user related to a sheet",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				/* res, err := server.RequireAuth(
					params.Context.Value("jwt").(string),
					QueryTodos,
				) */

				userID := params.Source.(model.Sheet).UserID
				res, err := users.QueryUserById(userID)

				if err != nil {
					return nil, err
				}

				return res, err
			},
		},
	},
})

var SegmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Segment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"sheetId": &graphql.Field{
			Type: graphql.Int,
		},
		"label": &graphql.Field{
			Type: graphql.String,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"createdAt": &graphql.Field{
			Type: graphql.String,
		},
	},
})
