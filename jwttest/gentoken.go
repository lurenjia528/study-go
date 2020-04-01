package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

var secret = "this is secret"

func main() {
	token, err := GenerateToken()
	if err != nil {
		panic(err)
	}
	println(token)
	valid(token)
}
func GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": "tom",
		//"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func valid(tokenStr string){
	token, _ := jwt.Parse(tokenStr,func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("not authorized")
			return nil, fmt.Errorf("not authorization")
		}
		return []byte(secret), nil
	})
	if !token.Valid {
		fmt.Println("not authorized")
	} else {
		fmt.Println("ok")
		println(token.Header)
	}
}