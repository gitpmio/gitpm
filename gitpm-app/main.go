package main

import (
	"github.com/astaxie/beego"
	_ "github.com/gitpmio/gitpm/gitpm-app/routers"
)

func main() {
	beego.Run()
}
