package main

import (
	//"errors"
	//"fmt"

	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"status": &graphql.Field{
			Type: graphql.Int,
		},

		// 直接关联书籍的查询
		// 怎么将这里的参数传递过去？
		"books": booksQuery,
	},
})
