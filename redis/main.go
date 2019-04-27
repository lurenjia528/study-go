package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

func main() {

	conn, err := redis.Dial("tcp", "localhost:6379")
	checkErr(err)
	defer conn.Close()

	//write(conn)
	//expire(conn)
	//exist(conn)
	//mutiRW(conn)
	//del(conn)
	//jsonRW(conn)
	listRW(conn)
}

// list
func listRW(conn redis.Conn){
	_, err := conn.Do("LPUSH", "username1", "zhangsan")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	_, err = conn.Do("LPUSH", "username1", "lisi")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}
	_, err = conn.Do("LPUSH", "username1", "wangwu")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	values, _ := redis.Values(conn.Do("LRANGE", "username1", "0", "2"))
	fmt.Printf("count=%d\n", len(values))
	for _, v := range values {
		fmt.Println(string(v.([]byte)))
	}
}

// 读写json
func jsonRW(conn redis.Conn) {
	cz := map[string]string{"name": "chengzi", "age": "19"}
	value, err := json.Marshal(cz)
	checkErr(err)
	reply, err := conn.Do("SETNX", "jsonkey", value)
	checkErr(err)
	if reply == int64(1) {
		fmt.Println("成功")
	}

	var result map[string]string
	bytes, err := redis.Bytes(conn.Do("GET", "jsonkey"))
	checkErr(err)
	json.Unmarshal(bytes,&result)
	fmt.Println(result["name"])
	fmt.Println(result["age"])
}

// 删除
func del(conn redis.Conn) {
	_, err := conn.Do("DEL", "qwe")
	checkErr(err)
}

// 批量读写
func mutiRW(conn redis.Conn) {
	_, err := conn.Do("MSET", "key1", "val1", "key2", 10)
	checkErr(err)
	s, err := redis.String(conn.Do("GET", "key1"))
	checkErr(err)
	fmt.Println(s)
	s1, err := redis.Int(conn.Do("GET", "key2"))
	checkErr(err)
	fmt.Println(s1)

	var value1 string
	var value2 int
	reply, err := redis.Values(conn.Do("MGET", "key1", "key2"))
	checkErr(err)
	_, err = redis.Scan(reply, &value1, &value2)
	checkErr(err)
	fmt.Println(value1)
	fmt.Println(value2)
}

// 判断是否存在
func exist(conn redis.Conn) {
	reply, err := conn.Do("EXISTS", "username")
	checkErr(err)
	if 1 == reply {
		fmt.Println("存在")
	}
	if 0 == reply {
		fmt.Println("不存在")
	}
}

// 读写
func write(conn redis.Conn) {
	_, err := conn.Do("SET", "username", "chengzi")
	checkErr(err)
	username, _ := redis.String(conn.Do("GET", "username"))
	fmt.Println(username)
}

// 写入10秒后过期
func expire(conn redis.Conn) {
	_, err := conn.Do("SET", "password", "123123", "EX", "10")
	checkErr(err)
	time.Sleep(11 * time.Second)
	pwd, err := redis.String(conn.Do("GET", "password"))
	checkErr(err)
	fmt.Println(pwd)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
