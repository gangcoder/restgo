package controllers

import (
	"github.com/astaxie/beego"
)

type MonitorController struct {
	beego.Controller
}

func (c *MonitorController) Get() {
	c.Ctx.WriteString("monitor")
	//c.TplName = "index.tpl"
}
