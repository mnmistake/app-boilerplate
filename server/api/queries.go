package api

import (
	"github.com/graphql-go/graphql"
	"github.com/raunofreiberg/kyrene/server"
	"github.com/raunofreiberg/kyrene/server/api/sheets"
	"github.com/raunofreiberg/kyrene/server/api/users"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"users": &graphql.Field{
			Type:        graphql.NewList(UserType),
			Description: "Query all users",
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
			Type:        SheetType,
			Description: "Query a sheet via its ID",
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type: graphql.Int,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				_, err := server.ValidateJWT(params.Context.Value("jwt").(string))

				if err != nil {
					return nil, err
				}

				sheetID, _ := params.Args["id"].(int)
				sheet, err := sheets.QuerySheet(sheetID)

				if err != nil {
					return nil, err
				}

				return sheet, nil
			},
		},
		"sheets": &graphql.Field{
			Type:        graphql.NewList(SheetType),
			Description: "Query all sheets",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				res, err := server.RequireAuth(
					params.Context.Value("jwt").(string),
					sheets.QuerySheets,
				)

				if err != nil {
					return nil, err
				}

				return res, err
			},
		},
	},
})
