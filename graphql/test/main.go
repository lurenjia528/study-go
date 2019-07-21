package main

import (
	"context"
	"fmt"
	"github.com/lurenjia528/study-go/graphql/test/model"
	"github.com/lurenjia528/study-go/graphql/test/object"
	"net/http"
	"strconv"

	"github.com/koding/multiconfig"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

type ServerCfg struct {
	Addr      string
	MysqlAddr string
}

func main() {
	//load config info
	m := multiconfig.NewWithPath("/home/lurenjia/gopath/src/github.com/lurenjia528/study-go/graphql/test/config.toml")
	svrCfg := new(ServerCfg)
	m.MustLoad(svrCfg)
	//new graphql schema
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    object.QueryType,
			Mutation: object.MutationType,
		},
	)
	if err != nil {
		fmt.Printf("[main] invoke graphql.NewSchema() failed,error: %s \n",err.Error())
		return
	}

	model.InitSqlxClient(svrCfg.MysqlAddr)
	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		//read user_id from gateway
		userIDStr := r.Header.Get("user_id")
		if len(userIDStr) > 0 {
			userID, err := strconv.Atoi(userIDStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(err.Error()))
				return
			}
			ctx = context.WithValue(ctx, "ContextUserIDKey", userID)
		}
		h.ContextHandler(ctx, w, r)

	})
	http.ListenAndServe(svrCfg.Addr, nil)
}
