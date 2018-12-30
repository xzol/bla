package controllers

import (
	"github.com/astaxie/beego"
)

type DocsController struct {
	beego.Controller
}
//
//func (c *DocsController) Get() {
//	c.Data["Website"] = "beego.me"
//	c.Data["Email"] = "astaxie@gmail.com"
//	c.TplName = "index.tpl"
//}

func (main *DocsController) HelloSitepoint() {
	main.Data["Website"] = "My Website"
	main.Data["Email"] = "your.email.address@example.com"
	main.Data["EmailName"] = "Your Name"
	main.Data["Id"] = main.Ctx.Input.Param(":id")
	main.TplName = "default/hello-sitepoint.tpl"
}
