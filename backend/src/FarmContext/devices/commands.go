package devices

type CreateDeviceCommand struct {
	DeviceSerial string
	Phone        string
	FarmId       uint
}

type AlertSetUpCommand struct {
	DeviceSerial string
	MinHumidity  uint
	MaxHumidity  uint
}
