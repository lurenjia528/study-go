// @APIVersion 1.0
// @Title 测试
// @Description
// @Contact
// @TermsOfServiceUrl
// @License
// @LicenseUrl
package routers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/lurenjia528/study-go/beego-study/hello/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/hw", &controllers.MainController{}, "get:Hw")

	beego.Get("/q", func(context *context.Context) {
		context.Output.Body([]byte("boo"))
	})

	beego.Post("/p", func(ctx *context.Context) {
		var re map[string]string
		json.Unmarshal(ctx.Input.RequestBody, &re)
		fmt.Println(re)
		//ctx.Output.Body([]byte("bob"))
	})

	//　自动路由
	beego.AutoRouter(&controllers.ObjectController{})

	// 注解路由
	beego.Include(&controllers.CMSController{})

	//	默认匹配 //例如对于URL"/api/123"可以匹配成功，此时变量":id"值为"123"
	//beego.Router("/api/?:id", &controllers.MainController{},"get:Test")
	//
	//	默认匹配 //例如对于URL"/api/123"可以匹配成功，此时变量":id"值为"123"，但URL"/api/"匹配失败
	//beego.Router("/api/:id", &controllers.MainController{},"get:Test")
	//
	//
	////	自定义正则匹配 //例如对于URL"/api/123"可以匹配成功，此时变量":id"值为"123"
	//beego.Router("/api/:id([0-9]+)", &controllers.MainController{},"get:Test")
	//
	//
	////	正则字符串匹配 //例如对于URL"/user/astaxie"可以匹配成功，此时变量":username"值为"astaxie"
	//beego.Router("/user/:username([\\w]+)", &controllers.MainController{},"get:Test")
	//
	//
	////*匹配方式 //例如对于URL"/download/file/api.xml"可以匹配成功，此时变量":path"值为"file/api"， ":ext"值为"xml"
	//beego.Router("/download/*.*", &controllers.MainController{},"get:Test")
	//
	//
	////*全匹配方式 //例如对于URL"/download/ceshi/file/api.json"可以匹配成功，此时变量":splat"值为"file/api.json"
	//beego.Router("/download/ceshi/*", &controllers.MainController{},"get:Test")
	//
	//
	////int 类型设置方式，匹配 :id为int 类型，框架帮你实现了正则 ([0-9]+)
	//beego.Router("/:id:int", &controllers.MainController{},"get:Test")
	//
	//
	////string 类型设置方式，匹配 :hi 为 string 类型。框架帮你实现了正则 ([\w]+)
	//beego.Router("/:hi:string/:age:int", &controllers.MainController{},"get:Test")
	//
	//
	////带有前缀的自定义正则 //匹配 :id 为正则类型。匹配 cms_123.html 这样的 url :id = 123
	//beego.Router("/cms_:id([0-9]+).html", &controllers.MainController{},"get:Test")

}
