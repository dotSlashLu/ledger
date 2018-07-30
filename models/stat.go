package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type ClassGroup struct {
	Count       int
	RecordCount int
	CostSum     float64
	Class       int
}

func StatGroupByClass(from time.Time) (*[]ClassGroup, error) {
	fromStr := from.Format("2006-01-02 15:04:05")
	o := orm.NewOrm()
	where := ""
	if !from.IsZero() {
		where += fmt.Sprintf(`
			WHERE create_time >= "%s"
		`, fromStr)
	}
	sql := fmt.Sprintf(`
		SELECT class, 
			COUNT(id) AS record_count, 
			SUM(cost) AS cost_sum 
		FROM expense
		%s
		GROUP BY class
	`, where)
	fmt.Printf("exec sql %s\n", sql)
	var group []ClassGroup

	if _, err := o.Raw(sql).QueryRows(&group); err != nil {
		return nil, err
	}
	fmt.Printf("got group %+v\n", group)
	return &group, nil
}
