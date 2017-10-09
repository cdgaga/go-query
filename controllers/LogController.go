package controllers

import (
	"github.com/astaxie/beego"
)

type LogListController struct {
	beego.Controller
}

func (c *LogListController) Get() {
    c.Data["Website"] = "beego.me"
    c.Data["Email"]   = "astaxie@gmail.com"
    c.TplName         = "log-list.html"
}