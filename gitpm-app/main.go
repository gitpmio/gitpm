package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/gitpmio/gitpm/gitpm-app/routers"
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:data.db")
}

func main() {
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Errorf(err)
	}
	beego.Run()
}
