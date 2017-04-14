package controllers

import (
	"github.com/astaxie/beego"
)

type RequestController struct {
	beego.Controller
}

func (c *RequestController) Get() {
	c.Ctx.WriteString("request")
	//c.TplName = "index.tpl"
}
