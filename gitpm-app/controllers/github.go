package controllers

import (
	"github.com/astaxie/beego"
)

type GithubController struct {
	beego.Controller
}

type User struct {
	Username string `json:"username"`
}

func (this *GithubController) Get() {
	username := this.Ctx.Input.Param(":username")
	u := User{
		Username: "rachit",
	}
	c.Data["json"] = &u
	c.ServeJSON()
}
