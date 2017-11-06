package controllers

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"github.com/cdgaga/go-query/models"
)

func (c *AdminController) DbConfigList() {
	o := orm.NewOrm()

	var configs []models.DbConfig
	num, err := o.Raw(" SELECT * FROM db_config ").QueryRows(&configs)
	if err == nil {
		fmt.Println("user nums: ", num)
	}
	fmt.Println(configs)

	c.TplName         = "db-config/list.html"
}

func (c *AdminController) DbConfigEdit() {
	db := models.GetDb()
	type Result struct {
		UserId int
		UserName  string
	}
	var result Result

	db.Raw("SELECT user_id, user_name FROM user WHERE user_id = ?", 1).Scan(&result)
	fmt.Println(result)
	// db.Exec("update user set is_super_admin =2;")

	c.TplName         = "db-config/edit.html"
}
