package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

// CMS API
type CMSController struct {
	beego.Controller
}


func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @Title getStaticBlock
// @Description get all the staticblock by key
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object} models.ZDTCustomer.Customer
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /staticblock/:key:string [get]
func (this *CMSController) StaticBlock() {
	key := this.Ctx.Input.Param(":key")
	fmt.Println("key:",key)
	this.Ctx.Output.Body([]byte("注解路由１"))
}

// @router /all/:key:string [get]
func (this *CMSController) AllBlock() {
	key := this.Ctx.Input.Param(":key")
	fmt.Println("key:",key)
	this.Ctx.Output.Body([]byte("注解路由２"))
}
