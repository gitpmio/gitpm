package controllers

import (
	"github.com/astaxie/beego"
)

type AsanaController struct {
	beego.Controller
}

type Webook struct {
	Type string `json:"type"`
}

func (webhook *Webook) processWebhook() bool {
	return true
}

func (asana *AsanaController) RecieveWebook() {
	webhook := Webook{
		Type: asana.GetString("type"),
	}
	processed := webhook.processWebhook()

	asana.Data["json"] = map[string]bool{"success": processed}
	asana.ServeJSON()
}
