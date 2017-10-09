package controllers

import (
	"github.com/astaxie/beego"
)

type RoleListController struct {
	beego.Controller
}

func (c *RoleListController) Get() {
    c.Data["Website"] = "beego.me"
    c.Data["Email"]   = "astaxie@gmail.com"
    c.TplName         = "role-list.html"
}