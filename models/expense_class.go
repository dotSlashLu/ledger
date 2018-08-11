package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type ExpenseClass struct {
	Id         int       `orm:"column(id);auto" json:"id"`
	Name       string    `orm:"column(name);size(32)" json:"name"`
	Level      int8      `orm:"column(level)" json:"level"`
	Parent     int       `orm:"column(parent);null" json:"parent"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" json:"create_time"`
}

type RankedClass struct {
	Class    ExpenseClass   `json:"class"`
	Children []ExpenseClass `json:"children"`
}

func (t *ExpenseClass) TableName() string {
	return "expense_class"
}

func init() {
	orm.RegisterModel(new(ExpenseClass))
}

// AddExpenseClass insert a new ExpenseClass into database and returns
// last inserted Id on success.
func AddExpenseClass(m *ExpenseClass) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetExpenseClassById retrieves ExpenseClass by Id. Returns error if
// Id doesn't exist
func GetExpenseClassById(id int) (v *ExpenseClass, err error) {
	o := orm.NewOrm()
	v = &ExpenseClass{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllRankedExpenseClass() ([]interface{}, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ExpenseClass))
	var l []ExpenseClass
	if _, err := qs.All(&l); err != nil {
		return nil, err
	}
	// 最终想要的结构：
	// [{class: cls11, children: [...]}, {class: cls12, children: [...]}]
	// 先弄成map {1: {class: , children: []}} 再扁平化
	hashMap := make(map[int]*RankedClass)
	for _, c := range l {
		if c.Level == 0 {
			if hashMap[c.Id] == nil {
				class := new(RankedClass)
				class.Class = c
				hashMap[c.Id] = class
			} else {
				hashMap[c.Id].Class = c
			}
		} else {
			p := hashMap[c.Parent]
			if p == nil {
				p = new(RankedClass)
				hashMap[c.Parent] = p
			}
			p.Children = append(p.Children, c)
		}
	}
	ret := make([]interface{}, len(hashMap))
	i := 0
	for _, v := range hashMap {
		ret[i] = v
		i++
	}
	fmt.Println(ret)
	return ret, nil
}

// GetAllExpenseClass retrieves all ExpenseClass matches certain condition. Returns empty list if
// no records exist
func GetAllExpenseClass(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ExpenseClass))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ExpenseClass
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateExpenseClass updates ExpenseClass by Id and returns error if
// the record to be updated doesn't exist
func UpdateExpenseClassById(m *ExpenseClass) (err error) {
	o := orm.NewOrm()
	v := ExpenseClass{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteExpenseClass deletes ExpenseClass by Id and returns error if
// the record to be deleted doesn't exist
func DeleteExpenseClass(id int) (err error) {
	o := orm.NewOrm()
	v := ExpenseClass{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ExpenseClass{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
