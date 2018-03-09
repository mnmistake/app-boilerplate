package sheets

import "github.com/graphql-go/graphql"

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
	},
})
