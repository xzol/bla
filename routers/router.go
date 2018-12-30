package routers

import (
	"bla.su/controllers"
	"github.com/astaxie/beego"
)
//
//# PUT request
//curl -X PUT http://localhost:8080/hello-world/213
//
//# DELETE request
//curl -X DELETE http://localhost:8080/hello-world/213

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/hello-world/:id([0-9]+)", &controllers.MainController{}, "get,post:HelloSitepoint")
	beego.Router("/litra", &controllers.LitraController{})
	//beego.Router("/litra/:id([0-9]+)", &controllers.LitraController{}, "get:Ls")

}
