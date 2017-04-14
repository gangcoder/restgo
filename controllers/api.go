package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {
	c.TplName = "api/index.tpl"
}
