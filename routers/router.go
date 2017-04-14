package routers

import (
	"github.com/astaxie/beego"
	"github.com/ironxu/restgo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/app", &controllers.AppController{})
	beego.Router("/api", &controllers.ApiController{})
	beego.Router("/request", &controllers.RequestController{})
	beego.Router("/series", &controllers.SeriesController{})
	beego.Router("/monitor", &controllers.MonitorController{})
	beego.Router("/doc", &controllers.DocController{})
	beego.Router("/auth", &controllers.AuthController{})
}
