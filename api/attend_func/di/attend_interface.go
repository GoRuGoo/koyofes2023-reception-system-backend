package di

import receivestruct "api/ReceiveStruct"

type AttendReturnStruct struct {
	UID                    *string  `json:"uid"`
	Name                   *string  `json:"name"`
	Attends_first_day      *bool    `json:"attends_first_day"`
	Attends_second_day     *bool    `json:"attends_second_day"`
	Temperature_first_day  *float64 `json:"temperature_first_day"`
	Temperature_second_day *float64 `json:"temperature_second_day"`
}

type Attend_Interface interface {
	ExistsUIDUser() (AttendReturnStruct, error)
	SetTemperature(receivestruct.PutTemperatureBodyStruct) error
}
