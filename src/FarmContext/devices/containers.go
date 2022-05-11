package devices

import (
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	DeviceSerial string `gorm:"column:device_serial"`
	Phone        string `gorm:"column:phone"`
	FarmID       uint
	MinHumidity  uint
	MaxHumidity  uint
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
