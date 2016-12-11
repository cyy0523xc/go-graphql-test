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
	Args: GetControllerArgs(graphql.FieldConfigArgument{
		"userId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	}),
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var res = make([]Book, 0)
		var userId uint32

		fmt.Printf("%+v\n", params)
		if id, ok := params.Args["userId"].(int); ok {
			userId = uint32(id)
		} else if user, ok := params.Source.(User); ok {
			// 如果父级对象是user，则获取其user.Id
			userId = user.Id
		}

		// 分页参数
		ctrlParams := ParseControllerArgs(params.Args)

		var index uint = 0
		for _, book := range books {
			if index < ctrlParams.Offset {
				index++
				continue
			}

			if book.UserId == userId {
				res = append(res, book)
			} else if userId == 0 {
				res = append(res, book)
			}
			index++

			if ctrlParams.Limit > 0 && index > ctrlParams.Offset+ctrlParams.Limit {
				break
			}
		}

		return res, nil
	},
}
