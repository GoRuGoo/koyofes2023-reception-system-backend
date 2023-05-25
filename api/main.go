package main

import (
	"api/attend_func/handler"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

type TemperatureStruct struct {
	Uid      string  `json:"uid"`
	Day      int     `json:"day"`
	Bodytemp float64 `json:"bodytemp"`
}

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	db := openSQL("mysql", accessPoint, 10)
	defer db.Close()

	router.GET("/attend", func(c *gin.Context) {
		uid := c.Query("uid")
		day, err := strconv.Atoi(c.Query("day"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errorMessage": "dayのstring->int変換に失敗"})
		}
		attendStructInstance := &handler.AttendStruct{DB: db, UID: &uid, Day: &day}
		err = handler.HandleExists(attendStructInstance)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"message": "success!"})
	})

	router.POST("/set_temperature", func(c *gin.Context) {
		var tempMiddleInstance TemperatureStruct
		err := c.ShouldBindJSON(&tempMiddleInstance)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		}
		temperatureInstance := &handler.AttendStruct{DB: db, UID: &tempMiddleInstance.Uid, BodyTemp: &tempMiddleInstance.Bodytemp, Day: &tempMiddleInstance.Day}
		err = handler.HandleSetTemperature(temperatureInstance)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errorMessage": err.Error()})
		}
	})
	router.Run(":8080")
}
