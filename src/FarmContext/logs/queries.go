package logs

import "time"

type GetDataFrameByDeviceIdQuery struct {
	DeviceId uint
	StartDate time.Time
	EndDate time.Time
	Step int
}

type GetDataFrameBySerialQuery struct {
	DeviceSerial string
	StartDate time.Time
	EndDate time.Time
	Step int
}

type GetWaterLogByDeviceIdQuery struct {
	DeviceId uint
}

type GetWaterLogBySerialQuery struct {
	DeviceSerial string
}

type GetWaterLogByIdQuery struct {
	Id uint
}