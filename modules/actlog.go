package modules

import (
	"github.com/cdgaga/go-query/models"
	"time"
)

const (
	ACT_TYPE_LOGIN = "login"
)

func ActlogCreate(actType string, actMap map[string]string) {

	var userId = 1
	var actTime string = time.Now().Format("2006-01-02 15:04:05")

	var data models.ActLog
	data.UserId = userId
	data.ActType = actType
	data.ActTitle = actMap["title"]
	data.ActContent = actMap["content"]
	data.ActTime = actTime

	models.GetDb().Create(&data)

	return
}
