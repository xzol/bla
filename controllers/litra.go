package controllers

import (
	"github.com/astaxie/beego"
)

type LitraController struct {
	beego.Controller
}

func (litra *LitraController) getNeedData() {
	litra.Data["appName"] = "litra";
	litra.Data["keywords"] = "Литература по програмированию, литература по ИТ, програмирование, хакеры, код, сайты, создание сайтов, написание программ."
	litra.Data["description"] = "Литература предназначенная для людей увлекающися програмированием и ИТ в целом."
}

func (litra *LitraController) Get() {
	litra.getNeedData()
	litra.TplName = "layoute/default.tpl"
}

func (litra *LitraController) Ls() {
	litra.getNeedData()
	litra.TplName = "layoute/default.tpl"
}
