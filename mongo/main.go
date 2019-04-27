package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Goods struct {
	ID    string
	Name  string
	Price string
	Url   string
}

func (a *Goods) save() error {
	s, c := connect("goods")
	defer s.Close()
	return c.Insert(&a)
}

func (a Goods) all() ([]Goods, error) {
	s, c := connect("goods")
	defer s.Close()
	var group []Goods
	err := c.Find(nil).All(&group)
	return group, err
}

func (a *Goods) get(id string) (*mgo.Query, *mgo.Session) {
	s, c := connect("goods")
	return c.Find(bson.M{"id": id}), s
}

func (a Goods) delete() error {
	s, c := connect("goods")
	defer s.Close()
	return c.Remove(bson.M{"id": a.ID})
}

func (a *Goods) update() (*mgo.Query, *mgo.Session) {
	s, c := connect("goods")
	defer s.Close()
	c.Update(bson.M{"id": a.ID}, a)
	return a.get(a.ID)
}

func connect(cName string) (*mgo.Session, *mgo.Collection) {
	session, err := mgo.Dial("127.0.0.1:27017") //Mongodb's connection
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	//return a instantiated collect
	return session, session.DB("mongo").C(cName)
}

func main() {
	goods := Goods{"123", "情侣装", "39", "https://www.taobao.com"}

	goods.save()

	//result,s := goods.get("123")
	//defer s.Close()
	//var re Goods
	//result.One(&re)
	//fmt.Println(re.ID)
	//fmt.Println(re.Name)
	//fmt.Println(re.Price)
	//fmt.Println(re.Url)

	//goods.Price = "49"
	//goods.update()
	//result,s := goods.get("123")
	//defer s.Close()
	//var re Goods
	//result.One(&re)
	//fmt.Println(re.ID)
	//fmt.Println(re.Name)
	//fmt.Println(re.Price)
	//fmt.Println(re.Url)

	//all, _ := goods.all()
	//for _,v := range all {
	//	println(v.ID)
	//	println(v.Name)
	//	println(v.Price)
	//	println(v.Url)
	//}

	//goods.delete()
}
