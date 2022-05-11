package logs

import (
	"context"
	"errors"
	"farm/src/FarmContext/alerts"
	"farm/src/FarmContext/devices"
	"farm/src/infrastructure"
)

func (d DeviceLog) validateForCreateNewInstance() error {
	return nil
}

func (w WaterLog) validateForCreateNewInstance() error {
	if w.Volume <= 0 {
		return errors.New("water volume cant be less than 0")
	}
	return nil
}

//Probably not in correct place
func checkHumidity(humidity uint, deviceSerial string, ctx context.Context) error {
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)
	alertRepo := alerts.NewAlertRepository(infrastructure.PostgresDBProvider)
	dev, err := deviceRepo.FindBySerial(ctx, deviceSerial)
	if err != nil {
		return err
	}
	if humidity > dev.MaxHumidity || humidity < dev.MinHumidity {
		alert, errr := alerts.NewAlert(
			alerts.CreateAlertParameters{Text: "Humidity out of range", DeviceSerial: deviceSerial})
		alertRepo.Save(ctx, alert)
		if errr != nil {
			return errr
		}
	}
	return nil
}
