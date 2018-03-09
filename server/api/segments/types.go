package segments

import "github.com/graphql-go/graphql"

var SegmentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Segment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"sheetId": &graphql.Field{
			Type: graphql.String,
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
