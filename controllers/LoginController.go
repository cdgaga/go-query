package controllers

import (
	"github.com/astaxie/beego"
	"github.com/cdgaga/go-query/models"
	"github.com/cdgaga/go-query/modules"
	"github.com/cdgaga/go-query/utils"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Prepare() {
	c.Data["staticUri"] = beego.AppConfig.String("staticUri")
	c.Data["Lang"] = "zh-CN"
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

// 登陆
func (c *LoginController) SignIn() {

	email := c.GetString("email")
	password := c.GetString("password")

	var user models.User
	models.GetDb().Where("user_name = ?", email).Find(&user)

	password = utils.StringMd5(password)
	mystruct := map[string]string{"code": "200", "message": ""}
	if password != user.UserPwd {
		mystruct["code"] = "500"
		mystruct["message"] = "password is error"
		c.Data["json"] = &mystruct
		c.ServeJSON()
	} else {

		c.SetSession("userName", user.UserName)
		c.SetSession("userId", user.UserId)

		modules.ActlogCreate(modules.ACT_TYPE_LOGIN, map[string]string{"title": "登录", "content": "登录成功"})

		url := beego.URLFor("AdminController.Home")
		mystruct["url"] = url
		c.Data["json"] = &mystruct
		c.ServeJSON()
	}
}

// 登出
func (c *LoginController) SignOut() {
	c.DestroySession()
	c.Redirect("/login", 302)
}
