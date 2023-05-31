package receivestruct

type PutTemperatureBodyStruct struct {
	BodyTempDay1 float64 `json:"temperature_day1"`
	BodyTempDay2 float64 `json:"temperature_day2"`
}
