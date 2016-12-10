package main

import (
	"github.com/graphql-go/graphql"
)

var commentType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Comment",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"content": &graphql.Field{
			Type: graphql.String,
		},
		"userId": &graphql.Field{
			Type: graphql.Int,
		},
		"bookId": &graphql.Field{
			Type: graphql.Int,
		},
	},
})
