package handler

import (
	"api/attend_func/di"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"
)

type ConnectBaseInfoStruct struct {
	DB  *sql.DB
	UID *string
}

// 時間判定用関数
func checkAvailableDateTime(dayInformation string) error {

	envVariable := fmt.Sprintf("DAY_%d_DATETIME", map[string]int{"temperature_first_day": 1, "temperature_second_day": 2}[dayInformation])

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
		return errors.New("日時が一致しません")
	}
	return nil
}

func (a *ConnectBaseInfoStruct) GetUserInfo() (di.ReturnAttendUserInfoStruct, error) {
	ReturnInstance := di.ReturnAttendUserInfoStruct{}

	err := a.DB.QueryRow("SELECT uid,name,attends_first_day,attends_second_day,temperature_first_day,temperature_second_day FROM reception WHERE uid = ?", *a.UID).Scan(&ReturnInstance.UID, &ReturnInstance.Name, &ReturnInstance.Attends_first_day, &ReturnInstance.Attends_second_day, &ReturnInstance.Temperature_first_day, &ReturnInstance.Temperature_second_day)
	if err != nil {
		return di.ReturnAttendUserInfoStruct{}, errors.New(err.Error())
	}
	return ReturnInstance, nil
}

func (a *ConnectBaseInfoStruct) SetTemperature(p di.ReceiveBodyTemperatureStruct) error {
	if p.BodyTempDay1 != 0 && p.BodyTempDay2 != 0 {
		return errors.New("体温が二つ入力されています。体温情報は改変出来ません。")
	} else if (p.BodyTempDay1 == 0 && p.BodyTempDay2 != 0) || (p.BodyTempDay1 != 0 && p.BodyTempDay2 == 0) {
		return errors.New("体温キーが二つ入力されています。片方にしてください。")
	} else if p.BodyTempDay1 == 0 && p.BodyTempDay2 == 0 {
		return errors.New("体温情報が空です。")
	}

	targetDay := "temperature_first_day"
	temp := p.BodyTempDay1
	if p.BodyTempDay2 != 0 {
		temp = p.BodyTempDay2
		targetDay = "temperature_second_day"
	}

	err := checkAvailableDateTime(targetDay)
	if err != nil {
		return errors.New(err.Error())
	}

	_, err = a.DB.Exec("UPDATE reception SET "+targetDay+" = ? WHERE uid = ?", temp, *a.UID)
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}

// handle function
func HandleGetUserInfo(a di.Attend_Interface) (di.ReturnAttendUserInfoStruct, error) {
	return a.GetUserInfo()
}

func HandleSetTemperature(a di.Attend_Interface, p di.ReceiveBodyTemperatureStruct) error {
	return a.SetTemperature(p)
}
