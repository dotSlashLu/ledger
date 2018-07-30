// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/dotSlashLu/ledger/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/expense",
			beego.NSInclude(
				&controllers.ExpenseController{},
			),
		),

		beego.NSNamespace("/expense_class",
			beego.NSInclude(
				&controllers.ExpenseClassController{},
			),
		),

		beego.NSNamespace("/stat",
			beego.NSGet("/class_group",
				controllers.StatController{}.GetClassGroup),
		),
	)
	beego.AddNamespace(ns)
}
