package controllers

import (
	"github.com/astaxie/beego"
)

type UserListController struct {
	beego.Controller
}

func (c *UserListController) Get() {
    c.Data["Website"] = "beego.me"
    c.Data["Email"]   = "astaxie@gmail.com"
    c.TplName         = "user-list.html"
}