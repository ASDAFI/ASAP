package logs

import (
	"gorm.io/gorm"
	"time"
)

type DeviceLog struct {
	DeviceSerial string    `gorm:"column:device_serial"`
	DeviceTime   time.Time `gorm:"column:datetime;" sql:"index:device_time_idx"`
	ServerTime   time.Time `gorm:"column:server_time"`
	Humidity     float32   `gorm:"column:humidity"`
}

type WaterLog struct {
	gorm.Model
	DeviceSerial string    `gorm:"column:device_serial"`
	UserId       uint      `gorm:"column:user_id"`
	Volume       float32   `gorm:"column:water_volume"`
	EntryTime    time.Time `gorm:"column:entry_time;" sql:"index:device_time_idx"`
}

type CreateDeviceLogParameters struct {
	DeviceSerial string
	DeviceTime   time.Time
	ServerTime   time.Time
	Humidity     float32
}

type CreateWaterLogParameters struct {
	DeviceSerial string
	UserId       uint
	Volume       float32
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

