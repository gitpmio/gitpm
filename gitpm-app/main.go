package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/gitpmio/gitpm/gitpm-app/routers"
	_ "github.com/mattn/go-sqlite3" // import your required driver
)

func init() {
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:data.db")
}

func main() {
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}
	beego.Run()
}
