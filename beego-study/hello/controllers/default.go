package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {

	jsoninfo := c.GetString("jsoninfo")
	fmt.Println(jsoninfo)
	asd := c.GetString("asd")
	fmt.Println(asd)
	// 接收文件
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		beego.Error("getfile err ", err)
	}
	defer f.Close()
	c.SaveToFile("uploadname", "static/upload/" + h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
}

func (c *MainController) Hw() {
	//c.EnableRender=false
	//c.Data["json"] = map[string]string{"ObjectId": "123"}
	//c.ServeJSON()
	s := beego.AppConfig.String("orange")
	fmt.Println(s)
	c.Ctx.WriteString("你好！")
}

//1
func (c *MainController) Prepare() {
	fmt.Println("prepare函数")
}

//2
func (c *MainController) Render() error {
	fmt.Println("render")
	return nil
}

//3
func (c *MainController) Finish() {
	fmt.Println("finish函数")
}

func (c *MainController) Test() {
	println("1-id:", c.Ctx.Input.Param(":id"))
	println("2-username:", c.Ctx.Input.Param(":username"))
	println("3-splat:", c.Ctx.Input.Param(":splat"))
	println("4-path:", c.Ctx.Input.Param(":path"))
	println("5-ext:", c.Ctx.Input.Param(":ext"))
	println("6-hi:", c.Ctx.Input.Param(":hi"))
	println("7-age:", c.Ctx.Input.Param(":age"))

}
