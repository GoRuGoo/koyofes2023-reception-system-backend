package main

import (
	"api/db"
	"api/models"
	"api/router"
)

func main() {
	db.Init()
	models.Init()
	models.InitReception()

	r := router.NewRouter()
	r.Run(":8080")
}
