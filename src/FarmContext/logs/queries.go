package logs

import "time"

type GetDataFrameByDeviceIdQuery struct {
	DeviceId  uint
	StartDate time.Time
	EndDate   time.Time
	Step      int
}

type GetDataFrameBySerialQuery struct {
	DeviceSerial string
	StartDate    time.Time
	EndDate      time.Time
	Step         int
}

type GetAllDataFramesByDeviceIdQuery struct {
	DeviceId uint
}

type GetAllDataFramesBySerialQuery struct {
	DeviceSerial string
}

type GetWaterLogBySerialQuery struct {
	DeviceSerial string
}

type GetWaterLogByIdQuery struct {
	Id uint
}

type GetWaterLogByFarmIdQuery struct {
	FarmId uint
}
