package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id          int
	Name        string
	Profile     *Profile   `orm:"rel(one)"` // OneToOne relation
	Post        []*Post `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Profile struct {
	Id          int
	Age         int16
	User        *User   `orm:"reverse(one)"` // 设置一对一反向关系(可选)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`    //设置一对多关系
	Tags  []*Tag `orm:"rel(m2m)"`
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init(){
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag))
	orm.RegisterDriver("mysql",orm.DRMySQL)
	orm.RegisterDataBase("default","mysql","root:123123@tcp/ygt?charset=utf8")
	orm.RunSyncdb("default", false, true)
	orm.RunSyncdb("default", false, true)
	orm.RunSyncdb("default", false, true)
	orm.RunSyncdb("default", false, true)
}

func main() {
	o := orm.NewOrm()

	o.Using("ygt")
	profile := new(Profile)
	profile.Age = 30

	user := new(User)
	user.Profile = profile
	user.Name = "slene"

	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))
}
