package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/graphql-go/graphql"
	//"github.com/graphql-go/handler"
)

// 用户
type User struct {
	Id     uint32 `json:"id"`
	Name   string `json:"name"`
	Status uint8  `json:"status"`
}

// 一个用户可以对应多本书
type Book struct {
	Id     uint32 `json:"id"`
	UserId uint32 `json:"userId"`
	Name   string `json:"name"`
}

// 一个用户可以对多本书进行评论，一本书也可以被多个用户评论
type Comment struct {
	Id      uint32 `json:"id"`
	UserId  uint32 `json:"userId"`
	BookId  uint32 `json:"bookId"`
	Content string `json:"content"`
}

var (
	users    []User
	books    []Book
	comments []Comment
)

// 跟查询定义
var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name:        "RootQuery",
	Description: "hello world",
	Fields: graphql.Fields{
		"user":  userQuery,
		"users": usersQuery,

		"book":  bookQuery,
		"books": booksQuery,

		"comment":  commentQuery,
		"comments": commentsQuery,

		"hello": &graphql.Field{
			Type:        graphql.String,
			Description: "hello world",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return "hello 中国人! ", nil
			},
		},
	},
})

// define schema, with our rootQuery and rootMutation
var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
	//Mutation: rootMutation,
})

func executeQuery(query string, paramsJson string, schema graphql.Schema) *graphql.Result {
	var params = make(map[string]interface{})
	if paramsJson != "" {
		if err := json.Unmarshal([]byte(paramsJson), &params); err != nil {
			fmt.Printf("params: %s, error: %s\n", paramsJson, err.Error())
		}
	}

	result := graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString:  query,
		VariableValues: params,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
func main() {
	url := "http://localhost:8080/graphql"
	fmt.Println("Now server is running on port 8080")
	fmt.Printf("查询接口详情：curl -g '%s?query={__type(name:\"User\"){name,description,fields{name,description,type{name}}}}'\n", url)
	fmt.Printf("获取单个用户：curl -g '%s?query={user(id:1){id,name}}'\n", url)
	fmt.Printf("获取单个用户：curl -g '%s?query=query+getUserList($id:Int){user(id:$id){id,name}}&params={\"id\":1}'\n", url)
	fmt.Printf("获取单个用户及其书名：curl -g '%s?query={user(id:1){id,name,books(userId:1){userId,name}}}'\n", url)

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.FormValue("query"), r.FormValue("params"), schema)
		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}
