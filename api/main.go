package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	accessPoint = "gorugo:gorupass@tcp(localhost:3306)/reception?parseTime=true%local=Asia%2FTokyo"
)

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.GET("/checkin", func(c *gin.Context) {
		uid := c.Query("uid")
		c.String(http.StatusOK, "hello %s uid", uid)
	})
	router.Run(":8080")
}
