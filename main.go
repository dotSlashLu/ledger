package main

import (
	_ "github.com/dotSlashLu/ledger/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	dbURI := beego.AppConfig.String("mysql_uri")
	orm.RegisterDataBase("default", "mysql", dbURI)

	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"*"},
	// 	AllowHeaders:     []string{"*"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	beego.Run()
}
