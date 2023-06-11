package controllers

import (
	"api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var usermodel = new(models.UsersModel)

func GetReceptionUserInfo(c *gin.Context) {
	uid := c.Param("uid")
	result, err := usermodel.GetReceptionUserInfo(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"name": result.Name, "uid": result.UID, "attends_day1": result.AttendsFirstDay, "attends_day2": result.AttendsSecondDay, "temperature_day1": result.TemperatureFirstDay, "temperature_day2": result.TemperatureSecondDay})
	return
}

func SetReceptionUserBodyTemperature(c *gin.Context) {
	uid := c.Param("uid")
	var receive_reception_user_body_temperature models.ReceiveReceptionUserBodyTemperature
	err := c.ShouldBindJSON(&receive_reception_user_body_temperature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = models.SetReceptionUserBodyTemperature(uid, receive_reception_user_body_temperature)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
