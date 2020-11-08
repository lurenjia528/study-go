package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/lurenjia528/study-go/graphql/monitor/cus"
	"github.com/lurenjia528/study-go/graphql/monitor/model"
	"net/http"
	"time"
	//"github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUser     string = "root"
	dbPassword string = "123123123"
	dbHost     string = "127.0.0.1"
	dbPort     int32  = 3306
	dbName     string = "graphql"
)

var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local&parseTime=true",
	dbUser, dbPassword, dbHost, dbPort, dbName)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func initDB() *gorm.DB {
	//获取连接
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//设置数据库连接池
	s, err := db.DB()
	if err != nil {
		panic(err)
	}
	s.SetMaxOpenConns(64)
	s.SetMaxIdleConns(16)
	s.SetConnMaxLifetime(time.Second * 300)

	if err = s.Ping(); err != nil {
		panic(err)
	}

	return db
}

func init() {
	db = initDB()
	//dsn := "root:123123123@tcp(127.0.0.1:3306)/graphql?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//if err != nil {
	//	fmt.Println("连接mysql失败", err)
	//}

	// 创建表
	err := db.AutoMigrate(&model.Monitor{})
	if err != nil {
		fmt.Println("创建表失败")
		panic(err)
	}
}


var monitor []model.Monitor
//var monitorS model.Monitor
var monitorType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Monitor",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: cus.Int64,
			},
			"deploymentName": &graphql.Field{
				Type: graphql.String,
			},
			"currentReplicas": &graphql.Field{
				Type: graphql.Int,
			},
			"availableReplicas": &graphql.Field{
				Type: graphql.Int,
			},
			"cpuLimit": &graphql.Field{
				Type: graphql.Int,
			},
			"cpuRequest": &graphql.Field{
				Type: graphql.Int,
			},
			"memLimit": &graphql.Field{
				Type: cus.Int64,
			},
			"memRequest": &graphql.Field{
				Type: cus.Int64,
			},
			"namespace": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"monitor": &graphql.Field{
				Type:        monitorType,
				Description: "查询监控信息",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(int)
					if ok {
						var monitorS model.Monitor
						db = DB()
						db.Find(&monitorS, id)
						return monitorS, nil
					}
					return nil, nil
				},
			},
			"list": &graphql.Field{
				Type:        graphql.NewList(monitorType),
				Description: "查询所有数据",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					//db = open()
					db = DB()
					db.Find(&monitor)
					return monitor, nil
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
		},
	)
	if len(result.Errors) > 0 {
		fmt.Printf("errors: %+v", result.Errors)
	}
	return result
}

func main() {
	http.HandleFunc("/monitor", func(w http.ResponseWriter, req *http.Request) {
		result := execQuery(req.URL.Query().Get("query"), schema)
		_ = json.NewEncoder(w).Encode(result)
	})
	fmt.Println("listen on 8080")
	_ = http.ListenAndServe(":8080", nil)

	//insert(db)
	//get(db)
	//update(db)
	//delete1(db)
}

//func insert(db *gorm.DB) {
//	monitor := model.Monitor{DeploymentName: "resource"}
//	db.Create(&monitor)
//}
//
//func get(db *gorm.DB) {
//	var monitor model.Monitor
//	db.First(&monitor)
//	fmt.Println(monitor)
//}
//
//func update(db *gorm.DB) {
//	var monitor model.Monitor
//	db.First(&monitor)
//	monitor.Namespace = "platform"
//	db.Save(&monitor)
//}
//
//func delete1(db *gorm.DB) {
//
//	db.Delete(&model.Monitor{}, "2")
//}
//func open() *gorm.DB {
//	if db != nil {
//		return db
//	}
//	var err error
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic(err)
//	}
//	return db
//}