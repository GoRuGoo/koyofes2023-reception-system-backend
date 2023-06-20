package main

import (
	"api/db"
	"api/models"
	"api/router"
	"fmt"
	"os"
)

func main() {
	db.Init()
	if os.Getenv("KOYOFES2023_MODE") == "DEBUG" {
		models.Init()
		models.InitReception()
		fmt.Println("Migration done.")
	}
	r := router.NewRouter()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}
