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

func parseISO8601(t string) (ret time.Time, err error) {
	if len(t) == 0 {
		ret = time.Time{}
	} else {
		ret, err = time.Parse(ISOLayout, t)
	}
	return
}

func parseRange(ctx *context.Context) (ok bool, from, to time.Time) {
	// from and to should be ISO 8601
	fromStr := ctx.Input.Query("from")
	toStr := ctx.Input.Query("to")

	from, err := parseISO8601(fromStr)
	if err != nil {
		returnError(ctx, err)
		return
	}
	to, err = parseISO8601(toStr)
	if err != nil {
		returnError(ctx, err)
		return
	}
	ok = true
	return
}

func (this StatController) ClassGroup(ctx *context.Context) {
	ok, from, to := parseRange(ctx)
	if !ok {
		return
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

func (this StatController) Overview(ctx *context.Context) {
	ok, from, to := parseRange(ctx)
	if !ok {
		return
	}

	overview, err := models.StatOverview(from, to)
	if err != nil {
		returnError(ctx, err)
		return
	} else {
		ctx.Output.SetStatus(200)
		ctx.Output.JSON(overview, true, false)
	}
}

func (this StatController) MonthGroup(ctx *context.Context) {
	group, err := models.StatGroupByMonth()
	if err != nil {
		returnError(ctx, err)
		return
	} else {
		ctx.Output.SetStatus(200)
		ctx.Output.JSON(group, true, false)
	}
}
