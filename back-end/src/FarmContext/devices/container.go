package devices

import (
	"errors"
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
	DeviceSerial string    `gorm:"column:device_serial"`
	Device       Device    `gorm:"-"`
	DeviceTime   time.Time `gorm:"column:datetime;" sql:"index:device_time_idx"`
	ServerTime   time.Time `gorm:"column:server_time"`
	Humidity     float32   `gorm:"column:humidity"`
	Temp         float32   `gorm:"column:temp"`
}

type WaterLog struct {
	gorm.Model
	DeviceId  uint      `gorm:"column:device_id"`
	UserId    uint      `gorm:"column:user_id"`
	Volume    float32   `gorm:"column:water_volume"`
	EntryTime time.Time `gorm:"column:entry_time;" sql:"index:device_time_idx"`
}

func NewWaterLog(deviceId uint, userId uint, volume float32, entryTime time.Time) (*WaterLog,error) {
	log :=  &WaterLog{DeviceId: deviceId, UserId: userId, Volume: volume, EntryTime: entryTime}

	return log, log.isValidForNew()
}

func (w WaterLog) isValidForNew() error {
	if w.Volume <= 0 {
		return errors.New("")
	}
	return nil
}

func (WaterLog) TableName() string {
	return "water_log"
}

func (Device) TableName() string {
	return "device"
}

func (DeviceLog) TableName() string {
	return "device_log"
}
