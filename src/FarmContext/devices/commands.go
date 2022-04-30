package devices

type CreateDeviceCommand struct {
	DeviceSerial string
	Phone        string
	FarmId       uint
}
