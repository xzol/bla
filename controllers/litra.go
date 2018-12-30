package controllers

import (
	"github.com/astaxie/beego"
)

type LitraController struct {
	beego.Controller
}

func (litra *LitraController) Get() {
	litra.Data["Website"] = "beego.me"
	litra.Data["Email"] = "astaxie@gmail.com"
	litra.TplName = "litra/ls.tpl"
}
