package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/gitpmio/gitpm/gitpm-app/models"
)

type UserController struct {
	beego.Controller
}

type Response struct {
	InsertedId int64  `json:"inserted_id"`
	Error      string `json:"error"`
}

func (c *UserController) Get() {
	u := models.User{Id: 1}
	user := u.Get()
	c.Data["json"] = map[string]string{
		"username": user.Username,
	}
	c.ServeJSON()
}

func (c *UserController) Post() {
	u := models.User{}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &u)

	id, err := u.Save()

	if err != nil {
		c.Data["json"] = Response{
			Error: err.Error(),
		}
	} else {
		c.Data["json"] = Response{
			InsertedId: id,
		}
	}
	c.ServeJSON()

}
