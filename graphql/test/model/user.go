package model

import (
	"database/sql"
	"fmt"
	"time"
)

type CommonInfo struct {
	ID          int            `db:"id"`
	CreatedTime time.Time      `db:"created_time"`
	DeletedTime sql.NullString `db:"deleted_time"`
	UpdatedTime sql.NullString `db:"updated_time"`
}

/*
请按照如下内容在数据库中建表
CREATE TABLE `t_user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `user_id` char(100) NOT NULL,
  `name` char(100) NOT NULL,
  `pwd` char(128) NOT NULL,
  `email` char(50) NOT NULL,
  `phone` char(20) NOT NULL,
  `status` int(10) NOT NULL,
  `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_time` char(100)  ,
  `updated_time` char(100)  ,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;
*/

type User struct {
	CommonInfo
	UserID string         `db:"user_id"`
	Name   string         `db:"name"`
	Pwd    string         `db:"pwd"`
	Email  sql.NullString `db:"email"`
	Phone  string         `db:"phone"`
	Status int64          `db:"status"`
}

func GetUser(userId string) (*User, error) {
	sql := `SELECT * FROM t_user where user_id = ? and deleted_time IS NOT null `
	var user User
	if err := dbx.Get(&user, sql, userId); err != nil {
		fmt.Printf("[model.GetUser] invoke mysql failed,error: %s \n", err.Error())
		return nil, err
	}
	return &user, nil
}

func GetUsers() ([]User, error) {
	var users []User
	sql := `SELECT * FROM t_user`
	if err := dbx.Select(&users, sql); err != nil {
		fmt.Printf("[model.GetUsers] invoke mysql failed,error: %s \n", err.Error())
		return nil, err
	}
	return users, nil

}

func InsertUser(user *User) error {
	sql := `INSERT into t_user (user_id, name, pwd, email, phone, status) VALUES (:user_id, :name, :pwd, :email, :phone, :status)`
	if _, err := dbx.NamedExec(sql, user); err != nil {
		fmt.Printf("[model.InserUser] invoke mysql failed,error: %s \n", err.Error())
		return err
	}
	return nil
}

//标记删除
func DeleteUser(userId string, status UserStatusType) error {
	sql := `UPDATE t_user SET status = :status , deleted_time = :deleted_time WHERE user_id = :user_id`
	user := User{
		UserID: userId,
		Status: int64(status),
		CommonInfo: CommonInfo{
			DeletedTime: struct {
				String string
				Valid  bool
			}{String: time.Now().Format("2006-01-02 15:04:05"), Valid: true},
		},
	}

	if _, err := dbx.NamedExec(sql, user); err != nil {
		fmt.Printf("[model.DeleteUser] invoke mysql failed,error: %s \n", err.Error())
		return err
	}
	return nil
}

func ChangeUserName(userId string, name string) error {
	sql := `UPDATE t_user SET name = :name , updated_time = :updated_time WHERE user_id = :user_id`
	user := User{
		UserID: userId,
		Name:   name,
		CommonInfo: CommonInfo{
			UpdatedTime: struct {
				String string
				Valid  bool
			}{String: time.Now().Format("2006-01-02 15:04:05"), Valid: true},
		},
	}
	if _, err := dbx.NamedExec(sql, user); err != nil {
		fmt.Printf("[model.ChangeUserName] invoke mysql failed,error: %s \n", err.Error())
		return err
	}
	return nil
}
