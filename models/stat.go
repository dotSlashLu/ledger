package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

const (
	dbTimeLayout = "2006-01-02 15:04:05"
	costDB       = "expense"
)

type ClassGroupStat struct {
	Count       int     `json:"count"`
	RecordCount int     `json:"record_count"`
	CostSum     float64 `json:"cost"`
	Class       int     `json:"class"`
}

type OverviewStat struct {
	CostDailyAvg float64 `json:"cost_daily_avg"`
	CostSum      float64 `json:"cost_sum"`
	TodaySum     float64 `json:"today_sum"`
}

type MonthGroupStat struct {
	CostSum float64 `json:"cost_sum"`
	Month   string  `json:"month"`
}

func timeRangeCriteria(criteria *[]string, from, to time.Time) {
	fromStr := from.Format(dbTimeLayout)
	toStr := to.Format(dbTimeLayout)
	if !from.IsZero() {
		*criteria = append(*criteria,
			fmt.Sprintf(`create_time >= "%s"`, fromStr))
	}
	if !to.IsZero() {
		*criteria = append(*criteria,
			fmt.Sprintf(`create_time < "%s"`, toStr))
	}
}

func StatGroupByClass(uid int, from, to time.Time) (*[]ClassGroupStat, error) {
	o := orm.NewOrm()

	criteria := []string{fmt.Sprintf(`uid = "%d"`, uid)}
	timeRangeCriteria(&criteria, from, to)

	where := ""
	if len(criteria) > 0 {
		where = "WHERE " + strings.Join(criteria, " AND ")
	}

	sql := fmt.Sprintf(`
		SELECT class,
			COUNT(id) AS record_count,
			SUM(cost) AS cost_sum
		FROM %s
		%s
		GROUP BY class
	`, costDB, where)
	fmt.Printf("exec sql %s\n", sql)
	var group []ClassGroupStat

	if _, err := o.Raw(sql).QueryRows(&group); err != nil {
		return nil, err
	}
	return &group, nil
}

func StatOverview(uid int, from, to time.Time) (*OverviewStat, error) {
	o := orm.NewOrm()

	criteria := []string{fmt.Sprintf(`uid = "%d"`, uid)}
	timeRangeCriteria(&criteria, from, to)

	where := ""
	if len(criteria) > 0 {
		where = "WHERE " + strings.Join(criteria, " AND ")
	}

	sql := fmt.Sprintf(`
		SELECT
			SUM(cost) / DATEDIFF(CURDATE(), MIN(create_time)) AS cost_daily_avg,
			SUM(cost) AS cost_sum,
			(SELECT SUM(cost) FROM %s WHERE
				create_time >= concat(CURDATE(), " 00:00:00") AND
				create_time <= concat(CURDATE(), " 23:59:59") AND
				uid = "%d") AS today_sum
		FROM %s
		%s
	`, costDB, uid, costDB, where)
	fmt.Printf("exec sql %s\n", sql)

	stat := new(OverviewStat)

	if err := o.Raw(sql).QueryRow(stat); err != nil {
		return nil, err
	}
	return stat, nil
}

func StatGroupByMonth(uid int, from, to time.Time) (*[]MonthGroupStat, error) {
	o := orm.NewOrm()
	criteria := []string{fmt.Sprintf(`uid = "%d"`, uid)}
	timeRangeCriteria(&criteria, from, to)
	where := ""
	if len(criteria) > 0 {
		where = "WHERE " + strings.Join(criteria, " AND ")
	}

	sql := fmt.Sprintf(`
		SELECT
			SUM(cost) AS cost_sum,
			MONTH(create_time) AS month
		FROM %s
		%s
		GROUP BY MONTH(create_time)
	`, costDB, where)
	fmt.Printf("exec sql %s\n", sql)
	group := new([]MonthGroupStat)
	if _, err := o.Raw(sql).QueryRows(group); err != nil {
		return nil, err
	}
	return group, nil
}
