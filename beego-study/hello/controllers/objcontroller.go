package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type ObjectController struct {
	beego.Controller
}

type Mystruct struct {
	Name string
	Age  int
}

func (c *ObjectController) Login() {
	c.EnableRender = false
	//c.Data["json"] = &Mystruct{"chengzi",20}
	//c.ServeJSON()
	//c.Data["xml"] = &Mystruct{"chengzi",20}
	//c.ServeXML()
	//c.Data["yaml"] = &Mystruct{"chengzi",20}
	//c.ServeYAML()
	fmt.Println("登录")
	beego.Emergency("this is emergency")
	beego.Alert("this is alert")
	beego.Critical("this is critical")
	beego.Error("this is error")
	beego.Warning("this is warning")
	beego.Notice("this is notice")
	beego.Informational("this is informational")
	beego.Debug("this is debug")

}
