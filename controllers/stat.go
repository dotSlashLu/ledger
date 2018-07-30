package controllers

import (
	"github.com/astaxie/beego/context"
	"github.com/dotSlashLu/ledger/models"
	"time"
)

const ISOLayout = "2006-01-02T15:04:05Z0700"

type StatController struct{}

type errReturn struct {
	Status  string
	Message string
}

func returnError(ctx *context.Context, err error) {
	ctx.Output.SetStatus(500)
	data := errReturn{"error", err.Error()}
	ctx.Output.JSON(data, true, false)
}

func (this StatController) GetClassGroup(ctx *context.Context) {
	// from and to should be ISO 8601
	fromStr := ctx.Input.Query("from")
	toStr := ctx.Input.Query("to")

	var from, to time.Time
	var err error
	if len(fromStr) == 0 {
		from = time.Time{}
	} else {
		from, err = time.Parse(ISOLayout, fromStr)
		if err != nil {
			returnError(ctx, err)
			return
		}
	}
	if len(toStr) == 0 {
		to = time.Time{}
	} else {
		to, err = time.Parse(ISOLayout, toStr)
		if err != nil {
			returnError(ctx, err)
			return
		}
	}
	groups, err := models.StatGroupByClass(from, to)
	if err != nil {
		returnError(ctx, err)
		return
	} else {
		ctx.Output.SetStatus(200)
		ctx.Output.JSON(groups, true, false)
	}
}
