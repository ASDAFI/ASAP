package controllers

import (
	"context"
	"farm/src/FarmContext/alerts"
	"farm/src/FarmContext/devices"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"
	pb_device "farm/src/proto/messages/device"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (f *FarmServer) GetAllAlerts(ctx context.Context, empty *emptypb.Empty) (*pb_device.Alerts, error) {

	userId := ctx.Value("user_id").(uint)
	log.Info("GetAllAlerts -- userId: ", userId)

	// Get User
	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQHandler := users.NewUserQueryHandler(userRepo)

	userQuery := users.GetUserByIdQuery{UserId: userId}
	user, err := userQHandler.GetUserById(ctx, userQuery)
	if err != nil {
		return nil, err
	}

	// Get Device
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)

	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)
	deviceQuery := devices.GetDeviceByFarmIdQuery{FarmId: user.FarmID}

	fetchedDevices, err := deviceQHandler.GetDeviceByFarmId(ctx, deviceQuery)
	if err != nil {
		return nil, err
	}

	alertRepo := alerts.NewAlertRepository(infrastructure.PostgresDBProvider)
	alertsQHandler := alerts.NewQueryHandler(alertRepo)

	result := make([]*pb_device.Alert, 0)
	for _, device := range fetchedDevices {
		alertQuery := alerts.FindAlertsByDeviceSerialQuery{DeviceSerial: device.DeviceSerial}
		fetchedAlerts, err := alertsQHandler.GetAlertsByDeviceSerial(ctx, alertQuery)
		if err != nil {
			return nil, err
		}
		for _, alert := range fetchedAlerts {
			result = append(result, &pb_device.Alert{
				Id:           uint32(alert.ID),
				DeviceSerial: alert.DeviceSerial,
				Text:         alert.Text,
				Humidity:     uint32(alert.Humidity),
			})
		}
	}

	response := &pb_device.Alerts{
		Alerts: result,
	}
	return response, nil
}

func (f *FarmServer) GetRecentAlerts(ctx context.Context, empty *emptypb.Empty) (*pb_device.Alerts, error) {

	userId := ctx.Value("user_id").(uint)
	log.Info("GetRecentAlerts -- userId: ", userId)

	// Get User
	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQHandler := users.NewUserQueryHandler(userRepo)

	userQuery := users.GetUserByIdQuery{UserId: userId}
	user, err := userQHandler.GetUserById(ctx, userQuery)
	if err != nil {
		return nil, err
	}

	// Get Device
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)

	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)
	deviceQuery := devices.GetDeviceByFarmIdQuery{FarmId: user.FarmID}

	fetchedDevices, err := deviceQHandler.GetDeviceByFarmId(ctx, deviceQuery)
	if err != nil {
		return nil, err
	}

	alertRepo := alerts.NewAlertRepository(infrastructure.PostgresDBProvider)
	alertsQHandler := alerts.NewQueryHandler(alertRepo)

	result := make([]*pb_device.Alert, 0)
	for _, device := range fetchedDevices {
		alertQuery := alerts.FindRecentAlertsByDeviceSerialQuery{DeviceSerial: device.DeviceSerial}
		fetchedAlerts, err := alertsQHandler.GetRecentAlertsByDeviceSerial(ctx, alertQuery)
		if err != nil {
			return nil, err
		}
		for _, alert := range fetchedAlerts {
			result = append(result, &pb_device.Alert{
				Id:           uint32(alert.ID),
				DeviceSerial: alert.DeviceSerial,
				Text:         alert.Text,
				Humidity:     uint32(alert.Humidity),
			})
		}
	}

	response := &pb_device.Alerts{
		Alerts: result,
	}
	return response, nil
}
