package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	db, err := sql.Open("postgres", "port=5432 user=postgres password=123123 dbname=ygt sslmode=disable")
	if err != nil {
		panic(err)
	}

	//Insert(db)
	Query(db)
	//Update(db)
	//Delete(db)
}


func Delete(db *sql.DB) {
	//删除数据
	stmt, err := db.Prepare("delete from userinfo where uid=$1")
	checkErr(err)

	res, err := stmt.Exec(1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}

func Update(db *sql.DB) {
	//更新数据
	stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	checkErr(err)

	res, err := stmt.Exec("ficow", 1)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("rows affect:", affect)
}

func Query(db *sql.DB) {
	rows, err := db.Query("SELECT uid,username,departname,created FROM userinfo")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string

		err := rows.Scan(&uid, &username, &department, &created)
		if err != nil {
			panic(err)
		}
		fmt.Println("uid = ", uid, "\nname = ", username, "\ndep = ", department, "\ncreated = ", created, "\n-----------")
	}
}

func Insert(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO userinfo (username,departname,created) VALUES ($1,$2,$3) RETURNING uid")
	if err != nil {
		panic(err)
	}

	result, err := stmt.Exec("chengzi", "生活部", "2019-10-01")
	if err != nil {
		panic(err)
	}

	affected, err := result.RowsAffected()
	checkErr(err)
	fmt.Println("row affect:", affected)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
