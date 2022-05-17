package alerts

import (
	"context"
	"farm/src/FarmContext/devices"
	"farm/src/FarmContext/farm"
	"farm/src/infrastructure"
)

func CheckAlert(ctx context.Context, deviceSerial string, humidity uint) error {
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)

	alertRepo := NewAlertRepository(infrastructure.PostgresDBProvider)

	alertCHandler := NewCommandHandler(alertRepo)

	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)

	device, err := deviceQHandler.GetDeviceBySerial(ctx, devices.GetDeviceBySerialQuery{deviceSerial})

	if err != nil {
		return err
	}

	if humidity < device.MinHumidity || humidity > device.MaxHumidity {
		alert, err := NewAlert(CreateAlertParameters{DeviceSerial: deviceSerial, Humidity: humidity, Text: "Out side the range"})
		if err != nil {
			return err
		}
		command := CreateAlertCommand{alert.DeviceSerial, alert.Text, alert.Humidity}

		alertCHandler.CreateAlert(ctx, command)

	}
	return nil

}
