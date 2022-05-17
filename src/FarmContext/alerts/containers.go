package alerts

import "gorm.io/gorm"

type Alert struct {
	gorm.Model
	DeviceSerial string `gorm:"column:device_serial"`
	Text         string
}

type CreateAlertParameters struct {
	DeviceSerial string
	Text         string
	Humidity     uint
}

func NewAlert(parameters CreateAlertParameters) (*Alert, error) {
	alert := &Alert{
		DeviceSerial: parameters.DeviceSerial,
		Text:         parameters.Text,
	}

	return alert, alert.validateForCreateNewInstance()
}
