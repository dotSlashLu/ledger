package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dotSlashLu/ledger/models"
)

type UserController struct {
	beego.Controller
}

func (this UserController) Status(ctx *context.Context) {
	type res struct {
		Login bool `json:"login"`
	}

	ret := res{}
	if this.Auth(ctx) {
		ret.Login = true
	}

	ctx.Output.SetStatus(200)
	ctx.Output.JSON(ret, true, false)
}

func (this UserController) Login(ctx *context.Context) {
	v := models.User{}
	if err := json.Unmarshal(ctx.Input.RequestBody, &v); err != nil {
		panic(err)
	}

	if len(v.Username) == 0 || len(v.Password) == 0 {
		ctx.Output.SetStatus(403)
		return
	}
	uid := models.CheckUser(v.Username, v.Password)
	if uid != 0 {
		ctx.Output.SetStatus(200)
		sess, _ := beego.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
		defer sess.SessionRelease(ctx.ResponseWriter)
		sess.Set("uid", uid)
		sess.Set("username", v.Username)
		ctx.Output.JSON(struct {
			Status string `json:"status"`
		}{"ok"}, true, false)
	}
}

func (this UserController) Logout(ctx *context.Context) {
	sess, _ := beego.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.ResponseWriter)

	sess.Delete("uid")
}

func (this UserController) Auth(ctx *context.Context) bool {
	// issue with session:
	// https://blog.csdn.net/qq_25504271/article/details/79411655
	// uid := this.GetSession("userid")
	sess, _ := beego.GlobalSessions.SessionStart(ctx.ResponseWriter, ctx.Request)
	defer sess.SessionRelease(ctx.ResponseWriter)
	uid := sess.Get("uid")
	if uid != nil {
		ctx.Input.SetData("uid", uid)
		return true
	}
	return false
}
