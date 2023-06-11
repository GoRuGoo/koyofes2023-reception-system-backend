package models

import (
	"api/db"
	"errors"
	"fmt"
	"os"
	"time"

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
type ReceiveReceptionUserBodyTemperature struct {
	BodyTempDay1 float32 `json:"temperature_day1"`
	BodyTempDay2 float32 `json:"temperature_day2"`
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

func SetReceptionUserBodyTemperature(uid string, r ReceiveReceptionUserBodyTemperature) error {
	if r.BodyTempDay1 != 0 && r.BodyTempDay2 != 0 {
		return errors.New("体温が二つ入力されています。")
	} else if r.BodyTempDay1 == 0 && r.BodyTempDay2 == 0 {
		return errors.New("体温情報が空です。")
	}

	targetDay := "temperature_first_day"
	temp := r.BodyTempDay1
	if r.BodyTempDay2 != 0 {
		targetDay = "temperature_second_day"
		temp = r.BodyTempDay2
	}

	err := isAcceptableTime(targetDay)
	if err != nil {
		return errors.New(err.Error())
	}

	testr := Reception{UID: uid}
	db.DB.Model(&testr).Update(targetDay, temp)
	return nil
}

func isAcceptableTime(targetDay string) error {
	envVariable := fmt.Sprintf("DAY_%d_DATETIME", map[string]int{"temperature_first_day": 1, "temperature_second_day": 2}[targetDay])
	envDateTimeStr := os.Getenv(envVariable)
	if envDateTimeStr == "" {
		return errors.New("時間に関する環境変数がセットされていません。")
	}

	envDateTime, err := time.Parse(time.RFC3339, envDateTimeStr)
	if err != nil {
		return errors.New(err.Error())
	}
	currentDateTime := time.Now()
	if currentDateTime.Year() != envDateTime.Year() || currentDateTime.YearDay() != envDateTime.YearDay() {
		return errors.New("受付可能日時ではありません。")
	}
	return nil
}
