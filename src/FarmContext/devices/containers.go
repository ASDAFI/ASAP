package devices

import (
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	DeviceSerial string `gorm:"column:device_serial;unique;not null"`
	Phone        string `gorm:"column:phone"`
	FarmID       uint   `gorm:"column:farm_id"`
	MinHumidity  uint   `gorm:"column:min_humidity"`
	MaxHumidity  uint   `gorm:"column:max_humidity"`
}

type CreateDeviceParameters struct {
	DeviceSerial string
	Phone        string
	FarmID       uint
}

func NewDevice(parameters CreateDeviceParameters) (*Device, error) {
	device := &Device{
		DeviceSerial: parameters.DeviceSerial,
		Phone:        parameters.Phone,
		FarmID:       parameters.FarmID,
		MinHumidity:  0,
		MaxHumidity:  100,
	}

	return device, device.validateForCreateNewInstance()
}

type SetUpAlertParameters struct {
	Device      *Device
	MinHumidity uint
	MaxHumidity uint
}

func SetUpAlert(parameters SetUpAlertParameters) error {
	parameters.Device.MinHumidity = parameters.MinHumidity
	parameters.Device.MaxHumidity = parameters.MaxHumidity
	return nil
}

func (Device) TableName() string {
	return "device"
}
