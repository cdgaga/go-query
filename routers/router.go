package routers

import (
	"github.com/cdgaga/go-query/controllers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"fmt"
)


func init() {

	// routing map
	beego.Router("/", &controllers.LoginController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/signIn", &controllers.LoginController{}, "post:SignIn")
	beego.Router("/signOut", &controllers.LoginController{}, "get:SignOut")
	beego.Router("/home", &controllers.AdminController{}, "get:Home")
	beego.Router("/profile", &controllers.AdminController{}, "get:Profile")

	beego.Router("/db/list", &controllers.AdminController{}, "get:DbList")
	beego.Router("/db/edit", &controllers.AdminController{}, "get:DbEdit")

	// user
	beego.Router("/user/list", &controllers.AdminController{}, "get:UserList")

	// loading language
	lang := "zh-CN"
	if err := i18n.SetMessage(lang, "conf/" + "locale_" + lang + ".ini"); err != nil {
		fmt.Println(err)
	}

	beego.AddFuncMap("i18n", i18n.Tr)
}

