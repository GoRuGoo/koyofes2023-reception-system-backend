package di

type Attend_Interface interface {
	ExistsUIDUser() (bool, error)
	SetTemperature() error
}
