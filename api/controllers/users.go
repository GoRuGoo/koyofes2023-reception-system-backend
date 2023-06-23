package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetReceptionUserInfo(c *gin.Context) {
	uid := c.Param("uid")
	result, err := models.GetReceptionUserInfo(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"name":             result.Name,
		"uid":              result.UID,
		"attends_day1":     result.AttendsFirstDay,
		"attends_day2":     result.AttendsSecondDay,
		"temperature_day1": result.TemperatureFirstDay,
		"temperature_day2": result.TemperatureSecondDay})
	return
}

func SetReceptionUserBodyTemperature(c *gin.Context) {
	uid := c.Param("uid")
	var receptionUserBodyTemperature models.ReceptionUserBodyTemperature
	err := c.ShouldBindJSON(&receptionUserBodyTemperature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.SetReceptionUserBodyTemperature(uid, receptionUserBodyTemperature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": receptionUserBodyTemperature})
	return
}
