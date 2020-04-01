package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/lurenjia528/study-go/beego-study/hello/controllers:CMSController"] = append(beego.GlobalControllerRouter["github.com/lurenjia528/study-go/beego-study/hello/controllers:CMSController"],
        beego.ControllerComments{
            Method: "AllBlock",
            Router: `/all/:key:string`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/lurenjia528/study-go/beego-study/hello/controllers:CMSController"] = append(beego.GlobalControllerRouter["github.com/lurenjia528/study-go/beego-study/hello/controllers:CMSController"],
        beego.ControllerComments{
            Method: "StaticBlock",
            Router: `/staticblock/:key:string`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
