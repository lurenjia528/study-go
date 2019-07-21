package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var dbx *sqlx.DB

func InitSqlxClient(connStr string) {
	var err error
	dbx, err = sqlx.Connect("mysql", connStr)
	if err != nil {
		fmt.Printf("[initSqlxClient] connect db:%s failed", connStr)
	}
}
