package main

import (
	_ "github.com/cdgaga/go-query/routers"
	"github.com/astaxie/beego"
	"github.com/cdgaga/go-query/controllers"
)

func main() {

	// open session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider="file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"

	// int err page
	beego.ErrorController(&controllers.ErrorController{})

	// run
	beego.Run()

}

