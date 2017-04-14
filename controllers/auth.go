package controllers

import (
	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Get() {
	c.Ctx.WriteString("auth")
	//c.TplName = "index.tpl"
}
