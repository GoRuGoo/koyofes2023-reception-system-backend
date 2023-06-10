package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := "root:gorupass@tcp(mysql:3306)/reception?charset=utf8mb4&parseTime=True"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
