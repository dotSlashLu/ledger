package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseClassController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Params:           nil})

	beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"] = append(beego.GlobalControllerRouter["github.com/dotSlashLu/ledger/controllers:ExpenseController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Params:           nil})

}
