package main

import (
	"github.com/graphql-go/graphql"
)

var commentQuery = &graphql.Field{
	Type: commentType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		idQuery, ok := params.Args["id"].(int)
		if ok {
			id := uint32(idQuery)
			for _, comment := range comments {
				if comment.Id == id {
					return comment, nil
				}
			}
		}

		return Comment{}, nil
	},
}

var commentsQuery = &graphql.Field{
	Type: graphql.NewList(commentType),
	Args: graphql.FieldConfigArgument{
		"userId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
		"bookId": &graphql.ArgumentConfig{
			Type: graphql.Int,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var res = make([]Comment, 0)
		var userOk, bookOk bool
		var userId, bookId uint32
		userIdQuery, userOk := params.Args["userId"].(uint32)
		bookIdQuery, bookOk := params.Args["bookId"].(uint32)
		userId, bookId = uint32(userIdQuery), uint32(bookIdQuery)

		for _, comment := range comments {
			if userOk {
				if bookOk {
					if comment.BookId == bookId && comment.UserId == userId {
						res = append(res, comment)
					}
				} else {
					if comment.UserId == userId {
						res = append(res, comment)
					}
				}
			} else if bookOk {
				if comment.BookId == bookId {
					res = append(res, comment)
				}
			}
		}

		return res, nil
	},
}
