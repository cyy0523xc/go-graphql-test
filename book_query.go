package main

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

var bookQuery = &graphql.Field{
	Type: bookType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		idQuery, ok := params.Args["id"].(int)
		if ok {
			id := uint32(idQuery)
			for _, book := range books {
				if book.Id == id {
					return book, nil
				}
			}
		}

		return Book{}, nil
	},
}

var booksQuery = &graphql.Field{
	Type: graphql.NewList(bookType),
	Args: graphql.FieldConfigArgument{
		"userId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var res = make([]Book, 0)
		userId, ok := params.Args["userId"].(int)
		fmt.Printf("%+v\n", params)
		if ok {
			id := uint32(userId)
			for _, book := range books {
				if book.UserId == id {
					res = append(res, book)
				}
			}
		}

		return res, nil
	},
}
