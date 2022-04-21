package devices

import "context"

type DeviceCommandHandler struct {
	waterLogRepository IWaterLogRepository
	deviceRepo         IDeviceRepository
}

func NewDeviceCommandHandler(waterLogRepository IWaterLogRepository, deviceRepo IDeviceRepository) *DeviceCommandHandler {
	return &DeviceCommandHandler{waterLogRepository: waterLogRepository, deviceRepo: deviceRepo}
}

func (h *DeviceCommandHandler) AddWaterHandler(ctx context.Context, command AddWaterCommand) error {
	device, err := h.deviceRepo.FindById(ctx, command.DeviceId)
	if err != nil {
		return err
	}
	waterLog, err := NewWaterLog(device.ID, command.UserId, command.Volume, command.Time)
	if err != nil {
		return err
	}

	if err := h.waterLogRepository.Save(ctx, waterLog); err != nil {
		return err
	}
	return err
}
func (h *DeviceCommandHandler) UpdateWaterHandler(ctx context.Context, command UpdateWaterCommand) error {
	device, err := h.deviceRepo.FindById(ctx, command.DeviceId)
	if err != nil {
		return err
	}
	waterLog, err := NewWaterLog(device.ID, command.UserId, command.Volume, command.Time)
	if err != nil {
		return err
	}

	if err := h.waterLogRepository.Updates(ctx, waterLog); err != nil {
		return err
	}
	return err
}
