package alerts

import "gorm.io/gorm"

type Alert struct {
	gorm.Model
	DeviceSerial string `gorm:"column:device_serial"`
	Text         string
	Humidity     uint
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
		Humidity:     parameters.Humidity,
	}

	return alert, alert.validateForCreateNewInstance()
}

func (Alert) TableName() string {
	return "alert"
}
