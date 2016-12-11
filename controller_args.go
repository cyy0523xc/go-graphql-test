package main

import (
	"github.com/graphql-go/graphql"
)

type CtrlArgs struct {
	Offset, Limit uint

	Sort string

	GroupBy string
}

var controllerArgs = graphql.FieldConfigArgument{
	"offset": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	"limit": &graphql.ArgumentConfig{
		Type: graphql.Int,
	},
	// 格式：fieldName1,-fieldName2
	// 默认升序，前面有-号表示降序
	"sort": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
	// 格式：fieldName1,fieldName2
	// 对应sql：group by fieldName1,fieldName2
	"groupby": &graphql.ArgumentConfig{
		Type: graphql.String,
	},
}

func GetControllerArgs(args graphql.FieldConfigArgument) graphql.FieldConfigArgument {
	for key, val := range controllerArgs {
		if _, ok := args[key]; ok {
			panic("GetControllerArgs for key: " + key)
		}
		args[key] = val
	}
	return args
}

func ParseControllerArgs(args map[string]interface{}) (ctrlParams *CtrlArgs) {
	ctrlParams = &CtrlArgs{}
	if val, ok := args["offset"]; ok {
		ctrlParams.Offset = uint(val.(int))
	}
	if val, ok := args["limit"]; ok {
		ctrlParams.Limit = uint(val.(int))
	}
	if val, ok := args["sort"]; ok {
		ctrlParams.Sort = val.(string)
	}
	if val, ok := args["groupby"]; ok {
		ctrlParams.GroupBy = val.(string)
	}

	return ctrlParams
}
