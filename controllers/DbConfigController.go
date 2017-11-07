package controllers

import (
	"fmt"
	"github.com/cdgaga/go-query/models"
	"github.com/cdgaga/go-query/utils"
	"strings"
)

func (c *AdminController) DbList() {

	offset := c.getOffset()

	dbTable := models.GetDb().Table("db_config")

	db_name := c.GetString("db_name")
	db_type := c.GetString("db_type")
	db_host := c.GetString("db_host")
	c.Data["db_name"] = db_name
	c.Data["db_type"] = db_type
	c.Data["db_host"] = db_host

	where := []string{}
	if db_name != "" {
		where = append(where, "db_name like '" + db_name + "%'")
	}

	if db_type != "" {
		where = append(where, "db_type like '" + db_type + "%'")
	}

	if db_host != "" {
		where = append(where, "db_host like '" + db_host + "%'")
	}

	whereStr := strings.Join(where, ",")
	if whereStr != "" {
		fmt.Println(whereStr)
		dbTable = dbTable.Where(whereStr)
	}


	var count int
	dbTable.Count(&count)

	var configs []models.DbConfig
	dbTable.Offset(offset).Limit(CONST_PAGE_SIZE).Find(&configs)

	p := utils.NewPaginator(c.Ctx.Request, CONST_PAGE_SIZE, count)
	c.Data["paginator"] = p
	c.Data["configs"] = configs
	c.TplName         = "db/list.html"
}

func (c *AdminController) DbEdit() {

	// id := c.GetString("id")
	c.TplName         = "db/edit.html"
}
