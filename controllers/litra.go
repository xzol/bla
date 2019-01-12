package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"io/ioutil"
	"log"
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
	//init
	litra.getNeedData()
	litra.TplName = "ajax/litra/ls.tpl"

	//data
	listFiles, err := ioutil.ReadDir("/home/x/K0d/go_projects/src/kod.rf/data/files")
	if err != nil {
		log.Fatal(err)
	}

	files := make(map[string][]string, 2)

	for _, file := range listFiles {
		if file.IsDir(){
			if files["d"] == nil{
				files["d"] = make([]string, 1)
			}
			if file.Name() == ""{
				files["d"] = append(files["d"], ".")
			}else {
				files["d"] = append(files["d"], file.Name())
			}
		}else{
			if files["-"] == nil {
				files["-"] = make([]string, 1)
			}
			files["-"] = append(files["-"], file.Name())
		}
	}

	// ********************* Marshal *********************
	b, err := json.Marshal(files)
	if err != nil {
		panic(err)
	}

	litra.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	litra.Ctx.ResponseWriter.Write(b)

}
