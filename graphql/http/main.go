package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"net/http"
)

type user struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var data map[string]user

/*
	创建user对象,拥有id,name属性
 */
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

/*
	创建查询类型
 */
var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					idQuery, isOk := p.Args["id"].(string)
					if isOk {
						return data[idQuery], nil
					}
					return nil, nil
				},
			},
		},
	},
)

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func execQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(
		graphql.Params{
			Schema:        schema,
			RequestString: query,
		})
	if len(result.Errors) > 0 {
		fmt.Printf("未知错误,error: %+v \n", result.Errors)
	}
	return result
}

func importJSONDataFromFile(fileName string, result interface{}) (isOk bool) {
	isOk = true
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		isOk = false
	}
	err = json.Unmarshal(content, result)
	if err != nil {
		isOk = false
		fmt.Println("Error:", err)
	}
	return
}

func main() {
	_ = importJSONDataFromFile("/home/lurenjia/gopath/src/github.com/lurenjia528/study-go/graphql/http/data.json", &data)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := execQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("listen on 8080")
	fmt.Println("测试地址:  curl -g 'http://127.0.0.1:8080/graphql?query={user(id:\"1\"){name}}'")
	http.ListenAndServe(":8080", nil)
}
