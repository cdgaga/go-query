package controllers

import (
	"github.com/astaxie/beego"
)



type AdminController struct {
	beego.Controller
}

// 常量 分页大小
const CONST_PAGE_SIZE int =  2

type PageStruct  struct {
	draw int // page
	recordsTotal int // page size
	recordsFiltered int //  page size
	data [][]string
}

func (c AdminController) getOffset() int {
	page, _ := c.GetInt("p")
	offset := 0;
	if page > 1 {
		offset =  (page - 1) * CONST_PAGE_SIZE
	}
	return  offset
}

func (c *AdminController) Prepare() {

	// 加载配置文件.
	c.Data["staticUri"] = beego.AppConfig.String("staticUri")

	// 检查用户是否登陆.
	userId := c.GetSession("userId")
	if userId == nil {
		c.Redirect("/login", 302)
	}

	c.Data["Lang"] = "zh-CN"
}
// 首页.
func (c *AdminController) Home() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"]   = "astaxie@gmail.com"
	c.TplName         = "home.html"
}

// 修过个人信息.
func (c *AdminController) Profile() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"]   = "astaxie@gmail.com"
	c.TplName         = "home.html"
}