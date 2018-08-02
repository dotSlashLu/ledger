package routers

import (
	"github.com/dotSlashLu/ledger/controllers"

	"github.com/astaxie/beego"
)

var statNS = beego.NSNamespace("/stat",
	beego.NSCond(controllers.UserController{}.Auth),
	beego.NSGet("/class_group",
		controllers.StatController{}.ClassGroup),
	beego.NSGet("/overview",
		controllers.StatController{}.Overview),
	beego.NSGet("/month_group",
		controllers.StatController{}.MonthGroup),
)
