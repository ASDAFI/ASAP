package devices

import (
	"context"
	"farm/src/FarmContext/farm"
)

type CommandHandler struct {
	deviceRepo IRepository
	farmRepo   farm.IRepository
}

func NewCommandHandler(deviceRepo IRepository, farmRepo farm.IRepository) *CommandHandler {
	return &CommandHandler{deviceRepo: deviceRepo, farmRepo: farmRepo}
}

func (h *CommandHandler) CreateDevice(ctx context.Context, command CreateDeviceCommand) error {
	farm, err := h.farmRepo.FindById(ctx, command.FarmId)
	if err != nil {
		return err
	}
	device, err := NewDevice(CreateDeviceParameters{DeviceSerial: command.DeviceSerial, Phone: command.Phone, FarmID: farm.ID})
	if err != nil {
		return err
	}

	if err := h.deviceRepo.Save(ctx, device); err != nil {
		return err
	}

	return nil

}

func (h *CommandHandler) SetUpAlert(ctx context.Context, command AlertSetUpCommand) error {
	dev, err := h.deviceRepo.FindBySerial(ctx, command.DeviceSerial)
	if err != nil {
		return err
	}
	err = SetUpAlert(SetUpAlertParameters{Device: dev, MinHumidity: command.MinHumidity,
		MaxHumidity: command.MaxHumidity})

	if err != nil {
		return err
	}
	err = h.deviceRepo.Save(ctx, dev)

	if err != nil {
		return err
	}
	return nil
}
