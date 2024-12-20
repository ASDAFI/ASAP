package devices

import "context"

type DeviceQueryHandler struct {
	deviceRepository IRepository
}

func NewDeviceQueryHandler(deviceRepository IRepository) *DeviceQueryHandler {
	return &DeviceQueryHandler{
		deviceRepository: deviceRepository,
	}
}

func (h *DeviceQueryHandler) GetDeviceById(ctx context.Context , query GetDeviceByIdQuery) (*Device, error) {
	device, err := h.deviceRepository.FindById(ctx, query.DeviceId)
	return device, err
}

func (h *DeviceQueryHandler) GetDeviceByFarmId(ctx context.Context, query GetDeviceByFarmIdQuery) ([]*Device, error) {
	devices, err := h.deviceRepository.FindByFarmId(ctx, query.FarmId)
	return devices, err
}

func (h *DeviceQueryHandler) GetDeviceBySerial(ctx context.Context, query GetDeviceBySerialQuery) (*Device, error) {
	device, err := h.deviceRepository.FindBySerial(ctx, query.DeviceSerial)
	return device, err
}