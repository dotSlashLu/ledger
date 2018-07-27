package main

import (
	_ "github.com/dotSlashLu/ledger/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	dbURI := beego.AppConfig.String("mysql_uri")
	orm.RegisterDataBase("default", "mysql", dbURI)
	beego.Run()
}
