package models

import "github.com/jinzhu/gorm"

type User struct {
	UserId    int
	UserName  string
	UserPwd   string
	Status        int
	LoginLastTime string
	LoginLastIp string
	CreateTime    string
	IsSuperAdmin  int
}

func (User) TableName() string {
	return "user"
}

type DbConfig struct {
	gorm.Model
	Id int
	DbName string
	DbType string
	DbHost string
	DbPassword string
	DbPort  string
	SubDataBases string
	CreateTime string
}
