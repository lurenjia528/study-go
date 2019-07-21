package object

import (
	"github.com/graphql-go/graphql"
	"github.com/lurenjia528/study-go/graphql/test/model"
)

type UserInfo struct {
	UserID string               `json:"userID"`
	Name   string               `json:"name"`
	Email  string               `json:"email"`
	Phone  string               `json:"phone"`
	Pwd    string               `json:"pwd"`
	Status model.UserStatusType `json:"status"`
}

var UserStatusEnumType = graphql.NewEnum(graphql.EnumConfig{
	Name:        "UserStatusEnum",
	Description: "用户状态信息",
	Values: graphql.EnumValueConfigMap{
		"EnableUser": &graphql.EnumValueConfig{
			Value:       model.EnableStatus,
			Description: "用户可用",
		},
		"DisableUser": &graphql.EnumValueConfig{
			Value:       model.DisableStatus,
			Description: "用户不可用",
		},
	},
})

var UserInfoType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "userInfo",
	Description: "用户信息描述",
	Fields: graphql.Fields{
		"userID": &graphql.Field{
			Description: "用户ID",
			Type:        graphql.String,
		},
		"name": &graphql.Field{
			Description: "用户名称",
			Type:        graphql.String,
		},
		"email": &graphql.Field{
			Description: "用户email",
			Type:        graphql.String,
		},
		"phone": &graphql.Field{
			Description: "用户手机号",
			Type:        graphql.String,
		},
		"pwd": &graphql.Field{
			Description: "用户密码",
			Type:        graphql.String,
		},
		"status": &graphql.Field{
			Description: "用户状态",
			Type:        UserStatusEnumType,
		},
	},
})
