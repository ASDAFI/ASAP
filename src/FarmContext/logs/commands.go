package logs

import "time"

type CreateDeviceLogCommand struct {
	DeviceSerial string
	DeviceTime   time.Time
	ServerTime   time.Time
	Humidity     uint
}

type CreateWaterLogCommand struct {
	DeviceSerial string
	UserId       uint
	Volume       float32
	EntryTime    time.Time
}
