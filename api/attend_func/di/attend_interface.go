package di

type ReturnAttendUserInfoStruct struct {
	UID                    *string  `json:"uid"`
	Name                   *string  `json:"name"`
	Attends_first_day      *bool    `json:"attends_first_day"`
	Attends_second_day     *bool    `json:"attends_second_day"`
	Temperature_first_day  *float64 `json:"temperature_first_day"`
	Temperature_second_day *float64 `json:"temperature_second_day"`
}

type ReceiveBodyTemperatureStruct struct {
	BodyTempDay1 float64 `json:"temperature_day1"`
	BodyTempDay2 float64 `json:"temperature_day2"`
}

type Attend_Interface interface {
	GetUserInfo() (ReturnAttendUserInfoStruct, error)
	SetTemperature(ReceiveBodyTemperatureStruct) error
}
