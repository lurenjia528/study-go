package main

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/beego/bee/logger"
)

func main() {
	result := httplib.Get("https://www.baidu.com")
	fmt.Println(result.String())

	beeLogger.Log.Info("info")
	beeLogger.Log.Success("success")
	beeLogger.Log.Warn("warn")
	beeLogger.Log.Critical("critical")
	beeLogger.Log.Fatal("fatal")


	//req := httplib.Post("https://www.baidu.com")
	//req.Param("key","value")
	//str, _ := req.String()
	//fmt.Println(str)

}
