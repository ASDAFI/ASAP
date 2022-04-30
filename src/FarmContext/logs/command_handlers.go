package logs

import (
	"context"
	"farm/src/FarmContext/devices"
)

type CommandHandler struct {
	deviceRepo devices.IRepository
	deviceLogRepo IDeviceLogRepository
	waterLogRepo IWaterLogRepository
}

func NewCommandHandler(deviceRepo devices.IRepository, deviceLogRepo IDeviceLogRepository, waterLogRepo IWaterLogRepository) *CommandHandler {
	return &CommandHandler{
		deviceRepo: deviceRepo,
		deviceLogRepo: deviceLogRepo,
		waterLogRepo:  waterLogRepo,
	}
}

func (h *CommandHandler) CreateDeviceLog(ctx context.Context, command CreateDeviceLogCommand) error {
	_, err := h.deviceRepo.FindBySerial(ctx, command.DeviceSerial)
	if err != nil {
		return err
	}
	deviceLog, err := NewDeviceLog(CreateDeviceLogParameters{DeviceSerial: command.DeviceSerial,
		                                                  DeviceTime: command.DeviceTime,
		                                                  ServerTime: command.ServerTime,
		                                                  Humidity: command.Humidity})
	if err != nil {
		return err
	}

	if err := h.deviceLogRepo.Save(ctx, deviceLog); err != nil{
		return err
	}

	return nil
}


func (h *CommandHandler) CreateWaterLog(ctx context.Context, command CreateWaterLogCommand) error {
	_, err := h.deviceRepo.FindBySerial(ctx, command.DeviceSerial)
	if err != nil {
		return err
	}
	waterLog, err := NewWaterLog(CreateWaterLogParameters{DeviceSerial: command.DeviceSerial,
		                                                  UserId: command.UserId,
		                                                  Volume: command.Volume,
		                                                  EntryTime: command.EntryTime,})
	if err != nil {
		return err
	}

	if err := h.waterLogRepo.Save(ctx, waterLog); err != nil{
		return err
	}

	return nil
}