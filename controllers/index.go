package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["appName"] = "index"
	c.Data["keywords"] = "Програмирование, хакеры, код, сайты, создание сайтов, написание программ."
	c.Data["description"] = "Портал предназначен для людей увлекающися програмированием и ИТ в целом."
	c.TplName = "layoute/default.tpl"
}

func (main *MainController) HelloSitepoint() {
	//main.Data["Website"] = "My Website"
	//main.Data["Email"] = "your.email.address@example.com"
	//main.Data["EmailName"] = "Your Name"
	//main.Data["Id"] = main.Ctx.Input.Param(":id")
	//main.TplName = "default/hello-sitepoint.tpl"
}
