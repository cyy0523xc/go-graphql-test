package main

import (
	//"errors"
	//"fmt"

	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "User",
	Description: "用户类型",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type:        graphql.Int,
			Description: "用户ID",
		},
		"name": &graphql.Field{
			Type:        graphql.String,
			Description: "用户名",
		},
		"status": &graphql.Field{
			Type:        userStatusType,
			Description: "用户状态",
		},

		"oldName": &graphql.Field{
			Type:              graphql.String,
			Description:       "旧书名",
			DeprecationReason: "已经过时的书名，弃用", // 前端查不到该字段了
		},

		// 直接关联书籍的查询
		// 怎么将这里的参数传递过去？
		"books": booksQuery,
	},
})

// 用户状态常量定义
const (
	UserStatusNormal uint8 = 0
	UserStatusMoney  uint8 = 1
	UserStatusHigh   uint8 = 2
)

// 接口中的用户类型定义
var userStatusType = graphql.NewEnum(graphql.EnumConfig{
	Name: "UserStatus",
	Values: graphql.EnumValueConfigMap{
		// 普通用户
		"NORMAL": &graphql.EnumValueConfig{
			Value: UserStatusNormal,
		},
		// 付费用户
		"MONEY": &graphql.EnumValueConfig{
			Value: UserStatusMoney,
		},
		// 高级用户
		"HIGH": &graphql.EnumValueConfig{
			Value: UserStatusHigh,
		},
	},
})
