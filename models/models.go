package models


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
	Id int
	DbName string
	DbType string
	DbHost string
	DbPassword string
	DbPort  string
	Charset  string
	SubDb string
	Tables string
	CreateTime string
}

func (DbConfig) TableName() string {
	return "db_config"
}