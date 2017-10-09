package routers

import (
	"github.com/cdgaga/go-query/controllers"
	"github.com/astaxie/beego"
)

func init() {

     beego.Router("/", &controllers.LoginController{})

     beego.Router("/home", &controllers.AdminController{}, "get:Home")
     beego.Router("/test/page", &controllers.AdminController{}, "get:TestPage")

     beego.Router("/db/query", &controllers.AdminController{}, "get:DbQuery")
     beego.Router("/db/query", &controllers.AdminController{}, "post:DbQueryResult")

}
