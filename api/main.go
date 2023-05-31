package main

import (
	"api/ReceiveStruct"
	"api/attend_func/handler"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	//	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

const (
	accessPoint = "gorugo:gorupass@tcp(localhost:3306)/reception"
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

type TemperatureStruct struct {
	UID      string  `json:"uid"`
	Day      int     `json:"day"`
	BodyTemp float64 `json:"bodytemp"`
}

func main() {
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
		attendStructInstance := &handler.AttendStruct{DB: db, UID: &uid}

		returnUserInfo, err := handler.HandleExists(attendStructInstance)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, returnUserInfo)
		return
	})

	router.PUT("/users/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		attendStructInstance := &handler.AttendStruct{DB: db, UID: &uid}
		var postedjson receivestruct.PutTemperatureBodyStruct
		err := c.ShouldBindJSON(&postedjson)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		err = handler.HandleSetTemperature(attendStructInstance, postedjson)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})

	//	router.POST("/set_temperature", func(c *gin.Context) {
	//		var tempMiddleInstance TemperatureStruct
	//		err := c.ShouldBindJSON(&tempMiddleInstance)
	//		if err != nil {
	//			c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
	//			return
	//		}
	//		temperatureInstance := &handler.AttendStruct{DB: db, UID: &tempMiddleInstance.UID, BodyTemp: &tempMiddleInstance.BodyTemp, Day: &tempMiddleInstance.Day}
	//		err = handler.HandleSetTemperature(temperatureInstance)
	//		if err != nil {
	//			c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
	//			return
	//		}
	//		c.JSON(http.StatusOK, gin.H{"message": "success!"})
	//		return
	//	})
	router.Run(":8080")
}
