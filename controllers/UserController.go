package controllers

import (
	"github.com/cdgaga/go-query/models"
	"github.com/cdgaga/go-query/utils"
)

func (c *AdminController) UserList()  {

	offset := c.getOffset()
	userName := c.GetString("user_name")

	dbTable := models.GetDb().Table("user")

	if userName != "" {
		dbTable =models.GetDb().Table("user").Where("user_name LIKE ?", "%"+userName+"%")
	}

	var count int
	dbTable.Count(&count)

	var users []models.User
	dbTable.Offset(offset).Limit(CONST_PAGE_SIZE).Find(&users);

	c.Data["users"] = users
	c.Data["user_name"] = userName

	p := utils.NewPaginator(c.Ctx.Request, CONST_PAGE_SIZE, count)
	c.Data["paginator"] = p
	c.TplName  = "user/list.html"
}
