package controllers

import (
	"context"

	"github.com/astaxie/beego"
	"github.com/google/go-github/github"
)

type GithubController struct {
	beego.Controller
}

type User struct {
	Organizations []*github.Organization `json:"username"`
}

func (this *GithubController) Get() {
	username := this.Ctx.Input.Param(":username")
	ctx := context.Background()
	client := github.NewClient(nil)
	// list all organizations for user "willnorris"
	orgs, _, _ := client.Organizations.List(ctx, username, nil)
	u := User{
		Organizations: orgs,
	}
	this.Data["json"] = &u
	this.ServeJSON()
}
