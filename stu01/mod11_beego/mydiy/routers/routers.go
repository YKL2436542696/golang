package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"myproject/controller"
)

func init() {
	beego.Router("/hello", &controller.MainController{})
}
