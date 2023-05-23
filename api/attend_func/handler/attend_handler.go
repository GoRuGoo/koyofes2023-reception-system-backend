package handler

import (
	"api/attend_func/di"
	"database/sql"
	"fmt"
)

type AttendStruct struct {
	DB       *sql.DB
	UID      *string
	BodyTemp float64
}

func (a *AttendStruct) ExistsUIDUser() (bool, error) {
	fmt.Println(*a.UID)
	return true, nil
}

func (a *AttendStruct) SetTemperature() error {
	return nil
}

func HandleExists(a di.Attend_Interface) {
	test, _ := a.ExistsUIDUser()
	fmt.Println(test)
}

func HandleSet(a di.Attend_Interface) {
	test := a.SetTemperature()
	fmt.Println(test)
}
