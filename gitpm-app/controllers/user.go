package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func (u *User) Save() bool {
	return true
}

func (c *UserController) Get() {
	u := User{
		Username: "rachit",
		Name:     "Rachit",
	}
	c.Data["json"] = &u
	c.ServeJSON()
}

func (c *UserController) Post() {
	username := c.GetString("username")
	name := c.GetString("name")

	u := User{
		Username: username,
		Name:     name,
	}

	saved := u.Save()

	c.Data["json"] = map[string]bool{"success": saved}
	c.ServeJSON()

}
