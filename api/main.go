package main

import (
	"api/attend_func/handler"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"

	"api/attend_func/di"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func openSQL(driverName, dataSourceName string, maxRetries int) (*sql.DB, error) {
	var db *sql.DB
	var err error
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open(driverName, dataSourceName)
		if err == nil {
			return db, nil
		}
		fmt.Println("DBとの接続に失敗", err)
		waitTime := time.Duration(i+1) * time.Second
		fmt.Println("再接続中", waitTime)
		time.Sleep(waitTime)
	}
	return nil, errors.New("DBとの接続に完全失敗")
}

func main() {

	accessPoint := "root:gorupass@tcp(mysql:3306)/reception"

	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	db, err := openSQL("mysql", accessPoint, 10)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	router.GET("/users/:uid", func(c *gin.Context) {

		uid := c.Param("uid")
		connectBaseInfo := &handler.ConnectBaseInfoStruct{DB: db, UID: &uid}

		forReturnUserInfo, err := handler.HandleGetUserInfo(connectBaseInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, forReturnUserInfo)
		return
	})

	router.PUT("/users/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		connectBaseInfo := &handler.ConnectBaseInfoStruct{DB: db, UID: &uid}
		var receiveUserBodyTemp di.ReceiveBodyTemperatureStruct
		err := c.ShouldBindJSON(&receiveUserBodyTemp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err = handler.HandleSetTemperature(connectBaseInfo, receiveUserBodyTemp)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})

	router.Run(":8080")
}
