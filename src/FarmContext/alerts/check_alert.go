package alerts

import (
	"context"
	"errors"
	"farm/src/FarmContext/devices"
	"farm/src/infrastructure"
	"time"
)

func CheckAlert(ctx context.Context, deviceSerial string, humidity uint) error {
	// get last alert
	alertRepo := NewAlertRepository(infrastructure.PostgresDBProvider)
	alertQHandler := NewQueryHandler(alertRepo)

	query := FindLastAlertByDeviceSerialQuery{
		DeviceSerial: deviceSerial,
	}

	lastAlert, err := alertQHandler.GetLastAlertByDeviceSerial(ctx, query)

	statement := false
	if err == nil {
		statement = lastAlert.CreatedAt.Unix()-time.Now().UTC().Unix() < 3600
	}

	if statement {
		return errors.New("Alert already sent")
	}

	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)

	alertCHandler := NewCommandHandler(alertRepo)

	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)

	device, err := deviceQHandler.GetDeviceBySerial(ctx, devices.GetDeviceBySerialQuery{DeviceSerial: deviceSerial})

	if err != nil {
		return err
	}

	if humidity < device.MinHumidity || humidity > device.MaxHumidity {
		alert, err := NewAlert(CreateAlertParameters{DeviceSerial: deviceSerial, Humidity: humidity, Text: "Out side the range"})
		if err != nil {
			return err
		}
		command := CreateAlertCommand{alert.DeviceSerial, alert.Text, alert.Humidity}

		return alertCHandler.CreateAlert(ctx, command)

	}
	return nil

}
