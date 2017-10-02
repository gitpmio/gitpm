package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/gitpmio/gitpm/gitpm-app/routers"
	_ "github.com/lib/pq" // import your required driver
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	postgresConnectionString := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s",
		beego.AppConfig.String("dbuser"),
		beego.AppConfig.String("dbpassword"),
		beego.AppConfig.String("dbhost"),
		beego.AppConfig.String("dbname"),
	)

	orm.RegisterDataBase("default", "postgres", postgresConnectionString)
}

func main() {
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}
	beego.Run()
}
