package handler

import (
	"api/attend_func/di"
	"database/sql"
	"errors"
	"fmt"
)

type AttendStruct struct {
	DB       *sql.DB
	UID      *string
	BodyTemp *float64
	Day      *int
}

func (a *AttendStruct) ExistsUIDUser() error {
	checkTargetDay := fmt.Sprintf("attends_%s_day", map[int]string{1: "first", 2: "second"}[*a.Day])

	var acceptance int = 0
	if *a.Day == 1 || *a.Day == 2 {

		err := a.DB.QueryRow("SELECT "+checkTargetDay+" FROM reception WHERE uid = ?", *a.UID).Scan(&acceptance)
		if err != nil {
			if err == sql.ErrNoRows {
				return errors.New("指定されたレコードが存在しません")
			}
			return errors.New("SELECTエラー")
		}

		fmt.Println(checkTargetDay)
		if acceptance > 0 {
			return nil
		} else {
			return errors.New("受付不可")
		}
	}
	return errors.New("1か2以外の数字がDayに入ってる")
}

func (a *AttendStruct) SetTemperature() error {
	tempTargetDay := fmt.Sprintf("temperature_%s_day", map[int]string{1: "first", 2: "second"}[*a.Day])

	err := a.ExistsUIDUser()
	if err != nil {
		return errors.New(err.Error())
	}

	if *a.Day == 1 || *a.Day == 2 {
		_, err := a.DB.Exec("UPDATE reception SET "+tempTargetDay+" = ? WHERE uid = ?", *a.BodyTemp, *a.UID)
		if err != nil {
			return errors.New(tempTargetDay + "の体温書き込み失敗")
		}
		return nil
	}
	return errors.New("1か2以外の数字がDayに入ってる")
}

func HandleExists(a di.Attend_Interface) error {
	err := a.ExistsUIDUser()
	return err
}

func HandleSetTemperature(a di.Attend_Interface) error {
	err := a.SetTemperature()
	return err
}
