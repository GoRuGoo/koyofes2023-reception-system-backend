package main

import (
	"api/attend_func/handler"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	accessPoint = "gorugo:gorupass@tcp(localhost:3306)/reception"
)

func openSQL(driverName, dataSourceName string, maxRetries int) *sql.DB {
	var db *sql.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open(driverName, dataSourceName)
		if err == nil {
			return db
		}
		fmt.Println("DBとの接続に失敗", err)
		waitTime := time.Duration(i+1) * time.Second
		fmt.Println("再接続中", waitTime)
		time.Sleep(waitTime)
	}
	fmt.Println("DBとの接続に完全失敗")
	os.Exit(1)
	return nil
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	db := openSQL("mysql", accessPoint, 10)
	var target_uid string = "33"
	attend_instance := &handler.AttendStruct{DB: db, UID: &target_uid}

	router.GET("/attend", func(c *gin.Context) {
		target_uid = c.Query("uid")
		handler.HandleExists(attend_instance)
	})
	//	router.GET("/checkin", func(c *gin.Context) {
	//		uid := c.Query("uid")
	//		c.String(http.StatusOK, "hello %s uid", uid)
	//	})
	router.Run(":8080")
}
