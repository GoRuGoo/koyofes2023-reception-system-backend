package models

import (
	"api/db"

	"gorm.io/gorm"
)

type Reception struct {
	UID                  string  `gorm:"primaryKey" json:"uid"`
	Mail                 string  `gorm:"not null" json:"mail"`
	Name                 string  `gorm:"not null" json:"name"`
	AttendsFirstDay      bool    `gorm:"not null" json:"attends_first_day"`
	AttendsSecondDay     bool    `gorm:"not null" json:"attends_second_day"`
	TemperatureFirstDay  float32 `json:"temperature_first_day"`
	TemperatureSecondDay float32 `json:"temperature_second_day"`
}

type UsersModel struct{}

func (u UsersModel) GetReceptionUserInfo(uid string) (Reception, error) {
	var result Reception
	var gorm_result *gorm.DB
	gorm_result = db.DB.First(&result, "uid = ?", uid)
	if err := gorm_result.Error; err != nil {
		return result, err
	}
	return result, nil
}
