package handler

import (
	"api/attend_func/di"
	"database/sql"
	"errors"
)

type AttendStruct struct {
	DB       *sql.DB
	UID      *string
	BodyTemp *float64
	Day      *int
}

func (a *AttendStruct) ExistsUIDUser() error {
	var acceptance bool = false
	if *a.Day == 1 {
		err := a.DB.QueryRow("SELECT attends_first_day FROM reception WHERE uid = ?", a.UID).Scan(&acceptance)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("指定されたレコードが存在しません")
			} else {
				return errors.New("SELECTエラー")
			}
		}
		return nil
	} else if *a.Day == 2 {
		err := a.DB.QueryRow("SELECT attends_second_day FROM reception WHERE uid = ?", a.UID).Scan(&acceptance)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("指定されたレコードが存在しません")
			} else {
				return errors.New("SELECTエラー")
			}
		}
		return nil
	} else {
		return errors.New("1か2以外の数字がDayに入ってる")
	}
}

func (a *AttendStruct) SetTemperature() error {
	if *a.Day == 1 {
		_, err := a.DB.Exec("UPDATE reception SET temperature_first_day = ? WHERE uid = ?", a.BodyTemp, a.UID)
		if err != nil {
			return errors.New("一日目の体温書き込み失敗")
		}
		return nil
	} else if *a.Day == 2 {
		_, err := a.DB.Exec("UPDATE reception SET temperature_second_day = ? WHERE uid = ?", a.BodyTemp, a.UID)
		if err != nil {
			return errors.New("二日目の体温の書き込み失敗")
		}
		return nil
	} else {
		return errors.New("1か2以外の数字がDayに入ってる")
	}
}

func HandleExists(a di.Attend_Interface) error {
	err := a.ExistsUIDUser()
	return err
}

func HandleSetTemperature(a di.Attend_Interface) error {
	err := a.SetTemperature()
	return err
}
