package routers

import (
	"github.com/astaxie/beego"
	"github.com/gitpmio/gitpm/gitpm-app/controllers"

	"github.com/gitpmio/gitpm/gitpm-app/controllers/asana"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/users", &controllers.UserController{})

	// Asana Routes
	beego.Router("/asana/webhook", &asana.AsanaController{}, "post:RecieveWebook")

	// example github Routes
	beego.Router("/github/:username", &controllers.GithubController{})
}
