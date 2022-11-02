package controller

import beego "github.com/beego/beego/v2/server/web"

type MainController struct {
	beego.Controller
}

// Get 姓名
func (c *MainController) Get() {
	c.Data["hello"] = "world"
}
