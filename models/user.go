package models

import (
	// "errors"
	// "fmt"
	// "reflect"
	// "strings"
	"time"

	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         int       `orm:"column(id);auto" json:"id"`
	Username   string    `orm:"column(username)" json:"username"`
	Password   string    `orm:"column(password)" json:"password"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);null" 
		json:"create_time"`
}

func init() {
	orm.RegisterModel(new(User))
}

func CheckUser(username, password string) int {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	if err := qs.Filter("username", username).One(user); err != nil {
		return 0
	}
	if !ComparePasswords(user.Password, password) {
		return 0
	}
	return user.Id
}

// AddExpense insert a new Expense into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func HashPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hash)
}

func ComparePasswords(hashed string, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}
