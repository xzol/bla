package controllers

import (
	"github.com/astaxie/beego"
)

type LitraController struct {
	beego.Controller
}

func (litra *LitraController) Get() {
	litra.Data["appName"] = "litra";
	litra.Data["keywords"] = "Литература по програмированию, литература по ИТ, програмирование, хакеры, код, сайты, создание сайтов, написание программ."
	litra.Data["description"] = "Литература предназначенная для людей увлекающися програмированием и ИТ в целом."
	litra.TplName = "layoute/default.tpl"
}
