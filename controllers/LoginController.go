package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {
    // 加载配置文件.
    c.Data["staticUri"] = beego.AppConfig.String("staticUri")
}

func (c *LoginController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"]   = "astaxie@gmail.com"
	c.TplName         = "login.html"
}
