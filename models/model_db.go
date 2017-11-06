package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/astaxie/beego"
)

var   gdb *gorm.DB

func init() {

	mysqlUser := beego.AppConfig.String("mysqlUser")
	mysqlPass := beego.AppConfig.String("mysqlPass")
	mysqlHost := beego.AppConfig.String("mysqlHost")
	mysqlPort := beego.AppConfig.String("mysqlPort")
	mysqlDb := beego.AppConfig.String("mysqlDb")

	connectStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlUser, mysqlPass, mysqlHost, mysqlPort, mysqlDb)
	// init Db
	Db, err := gorm.Open("mysql", connectStr)
	if err != nil {
		fmt.Println(err)
	}

	// enable dug
	Db.LogMode(true)

	gdb = Db

}

func GetDb() (*gorm.DB) {
	return gdb
}

func Close()  {
	gdb.Close()
}
