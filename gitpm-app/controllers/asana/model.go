package asana

import (
	"github.com/astaxie/beego/orm"
)

type Task struct {
	Id int
}

func init() {
	orm.RegisterModel(new(Task))
}

func (u *Task) Save() (id int64, err error) {
	o := orm.NewOrm()
	insertedId, err := o.Insert(u)

	if err != nil {
		return 0, err
	}

	return insertedId, nil
}

func (u *Task) Get() *Task {
	o := orm.NewOrm()
	o.Read(u)
	return u
}
