package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int
	Username string `json:"username,orm:"size(100)"`
	Name     string `json:"name",orm:"size(100)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) Save() (id int64, err error) {
	o := orm.NewOrm()
	insertedId, err := o.Insert(u)

	if err != nil {
		return 0, err
	}

	return insertedId, nil
}

func (u *User) Get() *User {
	o := orm.NewOrm()
	o.Read(u)
	return u
}
