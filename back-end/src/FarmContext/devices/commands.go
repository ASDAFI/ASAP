package devices

import "time"

type AddWaterCommand struct {
	DeviceId uint
	UserId uint
	Volume float32
	Time time.Time
}
