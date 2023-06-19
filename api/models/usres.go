package models

import (
	"api/db"
	"errors"
	"fmt"
	"os"
	"time"
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
type ReceptionUserBodyTemperature struct {
	BodyTempDay1 float32 `json:"temperature_day1"`
	BodyTempDay2 float32 `json:"temperature_day2"`
}

func GetReceptionUserInfo(uid string) (Reception, error) {
	var receptionUserInfo Reception
	err := db.DB.First(&receptionUserInfo, "uid = ?", uid).Error
	if err != nil {
		return receptionUserInfo, err
	}
	return receptionUserInfo, nil
}

func SetReceptionUserBodyTemperature(uid string, r ReceptionUserBodyTemperature) error {
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

	targetTable := Reception{UID: uid}
	err = db.DB.Model(&targetTable).Update(targetDay, temp).Error
	if err != nil {
		return err
	}
	return nil
}

func isAcceptableTime(targetDay string) error {
	//環境変数に設定された受付日時と現在時刻の年と月日のみを取り出して受付日時当日か検証する処理
	unnormalizedEnvReceptionDateTime := os.Getenv(fmt.Sprintf("DAY_%d_DATETIME", map[string]int{"temperature_first_day": 1, "temperature_second_day": 2}[targetDay]))
	if unnormalizedEnvReceptionDateTime == "" {
		return errors.New("時間に関する環境変数がセットされていません。")
	}

	//RFC3339の形式で正規化
	normalizedEnvDateTime, err := time.Parse(time.RFC3339, unnormalizedEnvReceptionDateTime)
	if err != nil {
		return errors.New(err.Error())
	}
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return errors.New(err.Error())
	}
	currentDateTime := time.Now().In(jst)

	if !(currentDateTime.After(normalizedEnvDateTime.Add(-time.Microsecond)) && currentDateTime.Before(normalizedEnvDateTime.AddDate(0, 0, 1))) {
		return errors.New("受付可能日ではありません。")
	}
	return nil
}
