package routers

import (
	"github.com/gitpmio/gitpm/gitpm-app/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
