package controllers

import (
	"github.com/astaxie/beego"
)

type SeriesController struct {
	beego.Controller
}

func (c *SeriesController) Get() {
	c.Ctx.WriteString("series")
	//c.TplName = "index.tpl"
}
