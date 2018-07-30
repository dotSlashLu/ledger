package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type ClassGroup struct {
	Count       int     `json:"count"`
	RecordCount int     `json:"record_count"`
	CostSum     float64 `json:"cost"`
	Class       int     `json:"class"`
}

const dbTimeLayout = "2006-01-02 15:04:05"

func StatGroupByClass(from, to time.Time) (*[]ClassGroup, error) {
	fromStr := from.Format(dbTimeLayout)
	toStr := to.Format(dbTimeLayout)
	o := orm.NewOrm()

	criteria := []string{}
	if !from.IsZero() {
		criteria = append(criteria,
			fmt.Sprintf(`create_time >= "%s"`, fromStr))
	}
	if !to.IsZero() {
		criteria = append(criteria,
			fmt.Sprintf(`create_time < "%s"`, toStr))
	}

	where := ""
	if len(criteria) > 0 {
		where = "WHERE " + strings.Join(criteria, " AND ")
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
	return &group, nil
}
