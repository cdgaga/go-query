package controllers

import (
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) Prepare() {

    // 加载配置文件.
    c.Data["staticUri"] = beego.AppConfig.String("staticUri")

    // todo 坚持用户是否登陆.
}

func (c *AdminController) Home() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"]   = "astaxie@gmail.com"
	c.TplName         = "home.html"
}

func (c *AdminController) TestPage() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"]   = "astaxie@gmail.com"
	c.TplName         = "test-page.html"
}