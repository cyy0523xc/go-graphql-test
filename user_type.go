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
			Type: userStatusType,
		},

		// 直接关联书籍的查询
		// 怎么将这里的参数传递过去？
		"books": booksQuery,
	},
})

var userStatusType = graphql.NewEnum(graphql.EnumConfig{
	Name: "UserStatus",
	Values: graphql.EnumValueConfigMap{
		// 普通用户
		"NORMAL": &graphql.EnumValueConfig{
			Value: uint8(0),
		},
		// 付费用户
		"MONEY": &graphql.EnumValueConfig{
			Value: uint8(1),
		},
		// 高级用户
		"HIGH": &graphql.EnumValueConfig{
			Value: uint8(2),
		},
	},
})
