package logs

import (
	"gorm.io/gorm"
	"time"
)

type DeviceLog struct {
	ID           uint      `gorm:"primarykey"`
	DeviceSerial string    `gorm:"column:device_serial"`
	DeviceTime   time.Time `gorm:"column:datetime;" sql:"index:device_time_idx"`
	ServerTime   time.Time `gorm:"column:server_time"`
	Humidity     uint      `gorm:"column:humidity"`
}

type WaterLog struct {
	gorm.Model
	DeviceSerial string    `gorm:"column:device_serial"`
	UserId       uint      `gorm:"column:user_id"`
	Volume       uint      `gorm:"column:water_volume"`
	EntryTime    time.Time `gorm:"column:entry_time;" sql:"index:device_time_idx"`
}

type CreateDeviceLogParameters struct {
	DeviceSerial string
	DeviceTime   time.Time
	ServerTime   time.Time
	Humidity     uint
}

type CreateWaterLogParameters struct {
	DeviceSerial string
	UserId       uint
	Volume       uint
	EntryTime    time.Time
}

func NewDeviceLog(params CreateDeviceLogParameters) (*DeviceLog, error) {
	deviceLog := &DeviceLog{
		DeviceSerial: params.DeviceSerial,
		DeviceTime:   params.DeviceTime,
		ServerTime:   params.ServerTime,
		Humidity:     params.Humidity,
	}
	return deviceLog, deviceLog.validateForCreateNewInstance()
}

func NewWaterLog(params CreateWaterLogParameters) (*WaterLog, error) {
	waterLog := &WaterLog{
		DeviceSerial: params.DeviceSerial,
		UserId:       params.UserId,
		Volume:       params.Volume,
		EntryTime:    params.EntryTime,
	}
	return waterLog, waterLog.validateForCreateNewInstance()
}

func (WaterLog) TableName() string {
	return "water_log"
}

func (DeviceLog) TableName() string {
	return "device_log"
}
