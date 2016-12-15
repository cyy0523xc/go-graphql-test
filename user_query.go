package main

import (
	"errors"

	"github.com/graphql-go/graphql"
)

var userQuery = &graphql.Field{
	Type: userType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		idQuery := params.Args["id"].(int)
		id := uint32(idQuery)
		for _, user := range users {
			if user.Id == id {
				return user, nil
			}
		}

		return nil, errors.New("no user")
	},
}

var usersQuery = &graphql.Field{
	Type: graphql.NewList(userType),
	Args: graphql.FieldConfigArgument{
		"status": &graphql.ArgumentConfig{
			Type:         userStatusType,
			DefaultValue: 0,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		var resUsers = make([]User, 0)
		statusQuery, ok := params.Args["status"].(int)
		println("status value: ", statusQuery)
		if ok {
			status := uint8(statusQuery)
			for _, user := range users {
				if user.Status == status {
					resUsers = append(resUsers, user)
				}
			}
			return resUsers, nil
		} else {
			return users, nil
		}

		return nil, errors.New("No user")
	},
}
