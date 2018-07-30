package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/context"
	"github.com/dotSlashLu/ledger/models"
	"time"
)

type StatController struct{}

func (this StatController) GetClassGroup(ctx *context.Context) {
	from := time.Now().Add(-time.Duration(3*24) * time.Hour)
	groups, err := models.StatGroupByClass(from)
	if err != nil {
		ctx.Output.SetStatus(500)
		ctx.Output.Body([]byte(err.Error()))
	} else {
		ctx.Output.SetStatus(200)
		text, _ := json.Marshal(groups)
		ctx.Output.Body(text)
	}
}
