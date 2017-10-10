package controllers

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)


func (c *AdminController) DbQuery() {

    c.TplName  = "db-query.html"
}

func (c *AdminController) DbQueryResult() {

    sql := c.GetString("sql")

    /*
    db, err := sql.Open("mysql", "root:pass@tcp(192.168.10.10:3306)/query_manage_tools?charset=utf8")
    if err != nil {
        fmt.Println("Open database error: %s\n", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        fmt.Println(err)
    }

    rows, err := db.Query("select * from qmt_user limit  ?", 3)
    if err != nil {
        fmt.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
    }*/
    orm.RegisterDataBase("default", "mysql", "root:pass@tcp(192.168.10.10:3306)/query_manage_tools?charset=utf8", 2)
    orm.Debug = true
    o := orm.NewOrm()

    var maps []orm.Params
    num, err := o.Raw(sql).Values(&maps)
    errMsg := ""
    

    if err == nil && num > 0 {
        
        var keys []string
        for key := range maps[0] {
            keys = append(keys, key)
        }

        fmt.Println(keys)

        c.Data["list"] = maps
        c.Data["keys"] = keys

        fmt.Println(keys)
    } else {
        fmt.Println(err)
        errMsg = err.Error()
    }

    // c.Data["sql"]    = sql
    c.Data["errMsg"] = errMsg

    

    c.TplName         = "db-query.html"
}