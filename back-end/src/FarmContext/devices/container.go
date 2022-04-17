package devices

import (
	"farm/src/FarmContext/farm"
	"gorm.io/gorm"
	"time"
)

type Device struct {
	gorm.Model
	DeviceSerial string    `gorm:"column:device_serial"`
	Phone        string    `gorm:"column:phone"`
	Farm         farm.Farm `gorm:"-"`
	FarmID       uint
}

type DeviceLog struct {
	DeviceSerial uint      `gorm:"column:device_serial"`
	Device       Device    `gorm:"-"`
	DeviceTime   time.Time `gorm:"column:datetime;" sql:"index:device_time_idx"`
	ServerTime   time.Time `gorm:"column:server_time"`
	Humidity     float32   `gorm:"column:humidity"`
	Temp         float32   `gorm:"column:temp"`
}

func (Device) TableName() string {
	return "device"
}

func (DeviceLog) TableName() string {
	return "device_log"
}
