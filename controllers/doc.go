package controllers

import (
	"github.com/astaxie/beego"
)

type DocController struct {
	beego.Controller
}

func (c *DocController) Get() {
	c.Ctx.WriteString("doc")
	//c.TplName = "index.tpl"
}
