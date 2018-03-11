package api

import (
	"github.com/graphql-go/graphql"
	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/api/segments"
	"github.com/raunofreiberg/kyrene/server/api/sheets"
	"github.com/raunofreiberg/kyrene/server/api/users"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(users.UserType),
			Description: "Return all users",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				res, err := server.RequireAuth(
					params.Context.Value("jwt").(string),
					users.QueryUsers,
				)

				if err != nil {
					return nil, err
				}

				return res, nil
			},
		},
		"sheet": &graphql.Field{
			Type:        sheets.SheetType,
			Description: "Query a sheet",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				sheetID, _ := params.Args["id"].(int)
				sheet, err := sheets.QuerySheet(sheetID)

				if err != nil {
					return nil, err
				}

				return sheet, nil
			},
		},
		"sheets": &graphql.Field{
			Type:        graphql.NewList(sheets.SheetType),
			Description: "Query all sheets",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				/* res, err := server.RequireAuth(
					params.Context.Value("jwt").(string),
					QueryTodos,
				) */

				res, err := sheets.QuerySheets()

				if err != nil {
					return nil, err
				}

				return res, err
			},
		},
		"segments": &graphql.Field{
			Type:        graphql.NewList(segments.SegmentType),
			Description: "Query all segments",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				/* res, err := server.RequireAuth(
					params.Context.Value("jwt").(string),
					QueryTodos,
				) */

				sheetID := params.Args["id"].(int)

				res, err := segments.QuerySegments(sheetID)

				if err != nil {
					return nil, err
				}

				return res, err
			},
		},
	},
})
