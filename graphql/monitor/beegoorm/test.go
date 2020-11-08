package beegoorm


import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/lurenjia528/study-go/graphql/monitor/model"
)

func test() {
	//database, err := sqlx.Open("mysql", "root:XXXX@tcp(127.0.0.1:3306)/test")
	//    //database, err := sqlx.Open("数据库类型", "用户名:密码@tcp(地址:端口)/数据库名")
	//1.连接数据库
	orm.RegisterDataBase("default", "mysql", "root:123123123@tcp(127.0.0.1:3306)/graphql")

	//2.注册表
	orm.RegisterModel(new(model.Monitor))

	//3.生成表
	orm.RunSyncdb("default",false,true)

	//doInsertOrm()
	queryOrm := doQueryOrm()
	//doUpdateOrm(queryOrm)
	doDeleteOrm(queryOrm)
}

func doInsertOrm() {
	fmt.Println("开始doQuerryOrm")

	//1.定义一个orm对象
	o := orm.NewOrm()
	var monitor model.Monitor

	//定义要插入的bean
	monitor.DeploymentName = "resource"

	id, err := o.Insert(&monitor)
	//utils.DoError("orm插入错误", err)
	if err == nil {
		fmt.Println("插入成功_id:", id)
	}
}

func doQueryOrm() (monitor *model.Monitor){
	o := orm.NewOrm()
	//查询操作
	var monitor1  model.Monitor
	//monitor.DeploymentName = "resource"
	_,err := o.QueryTable(monitor1).All(&monitor1)
	//err := o.Read(&monitor)
	if err != nil {
		fmt.Println(err)
	}
	//utils.DoError("orm查询错误", err)
	fmt.Println(monitor1)
	return &monitor1
}

func doUpdateOrm(monitor *model.Monitor) {
	o := orm.NewOrm()
	//更新操作  --需要先查询
	monitor.Namespace = "platform"
	count, err := o.Update(monitor)
	if err != nil {
		fmt.Println(err)
	}
	//utils.DoError("orm更新错误", err)
	fmt.Println("orm更新成功条目数量:", count)
}
func doDeleteOrm(monitor *model.Monitor) {
	o := orm.NewOrm()
	//更新操作  --需要先查询
	count, err := o.Delete(monitor)
	if err != nil {
		fmt.Println(err)
	}
	//utils.DoError("orm更新错误", err)
	fmt.Println("orm更新成功条目数量:", count)
}
