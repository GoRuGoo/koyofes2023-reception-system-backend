package models

import "api/db"

func Init() {
	db.DB.AutoMigrate(&Reception{})
}

func InitReception() {
	var users = []Reception{
		{UID: "33u@2", Mail: "goru.technology@gmail.com", Name: "伊藤優汰", AttendsFirstDay: true, AttendsSecondDay: false},
	}
	db.DB.Save(&users)
}
