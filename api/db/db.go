package db

import (
	"log"

	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Init() {
	var err error
	name := os.Getenv("PLANET_SCALE_USER_NAME")
	password := os.Getenv("PLANET_SCALE_USER_PASSWORD")
	ip := os.Getenv("PLANET_SCALE_IP")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/koyofes2023-reception?tls=True", name, password, ip)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
