package controllers

import (
	"github.com/astaxie/beego"
)

type AppController struct {
	beego.Controller
}

func (c *AppController) Get() {
	c.Ctx.WriteString("app")
	//c.TplName = "index.tpl"
}
