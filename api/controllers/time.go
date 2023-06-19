package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetEnvironmentTime(c *gin.Context) {
	day1_time := os.Getenv("DAY_1_DATETIME")
	day2_time := os.Getenv("DAY_2_DATETIME")

	if day1_time == "" || day2_time == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "環境変数が設定されていません。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"day1_envtime": day1_time, "day2_envtime": day2_time})
	return
}
