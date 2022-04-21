package devices

import "time"

type AddWaterCommand struct {
	DeviceId uint
	UserId uint
	Volume float32
	Time time.Time
}
type UpdateWaterCommand struct {
	Id uint
	DeviceId uint
	UserId uint
	Volume float32
	Time time.Time
}
