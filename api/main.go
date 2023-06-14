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
	models.Init()
	models.InitReception()

	r := router.NewRouter()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	r.Run(port)
}
