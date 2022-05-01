package logs

import (
	"context"
	"farm/src/FarmContext/devices"
	"farm/src/FarmContext/farm"
)

type LogQueryHandler struct {
	deviceLogRepository IDeviceLogRepository
	waterLogRepository  IWaterLogRepository
	deviceRepository    devices.IRepository
	farmRepository      farm.IRepository
}

func NewLogQueryHandler(deviceLogRepository IDeviceLogRepository, waterLogRepository IWaterLogRepository,
	deviceRepository devices.IRepository, farmRepository farm.IRepository) *LogQueryHandler {
	return &LogQueryHandler{
		deviceLogRepository: deviceLogRepository,
		waterLogRepository:  waterLogRepository,
		deviceRepository:    deviceRepository,
		farmRepository:      farmRepository,
	}
}

func (h *LogQueryHandler) GetDataFrameByDeviceId(ctx context.Context, query *GetDataFrameByDeviceIdQuery) ([]*DeviceLog, error) {
	_, err := h.deviceRepository.FindById(ctx, query.DeviceId)

	if err != nil {
		return nil, err
	}

	deviceLogs, err := h.deviceLogRepository.GetDataFrameByDeviceId(ctx, query.DeviceId, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	dataFrame := make([]*DeviceLog, 0)
	for i := 0; i < len(deviceLogs); i = i + query.Step {
		dataFrame = append(dataFrame, deviceLogs[i])
	}

	return dataFrame, nil
}

func (h *LogQueryHandler) GetDataFrameBySerial(ctx context.Context, query GetDataFrameBySerialQuery) ([]*DeviceLog, error) {
	_, err := h.deviceRepository.FindBySerial(ctx, query.DeviceSerial)

	if err != nil {
		return nil, err
	}

	deviceLogs, err := h.deviceLogRepository.GetDataFrameBySerial(ctx, query.DeviceSerial, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	dataFrame := make([]*DeviceLog, 0)
	for i := 0; i < len(deviceLogs); i = i + query.Step {
		dataFrame = append(dataFrame, deviceLogs[i])
	}

	return dataFrame, nil
}

func (h *LogQueryHandler) GetAllDataFramesByDeviceId(ctx context.Context, query *GetAllDataFramesByDeviceIdQuery) ([]*DeviceLog, error) {
	_, err := h.deviceRepository.FindById(ctx, query.DeviceId)

	if err != nil {
		return nil, err
	}

	deviceLogs, err := h.deviceLogRepository.GetAllDataFramesByDeviceId(ctx, query.DeviceId)
	if err != nil {
		return nil, err
	}

	return deviceLogs, nil
}

func (h *LogQueryHandler) GetAllDataFramesBySerial(ctx context.Context, query GetAllDataFramesBySerialQuery) ([]*DeviceLog, error) {
	_, err := h.deviceRepository.FindBySerial(ctx, query.DeviceSerial)

	if err != nil {
		return nil, err
	}

	deviceLogs, err := h.deviceLogRepository.GetAllDataFramesBySerial(ctx, query.DeviceSerial)
	if err != nil {
		return nil, err
	}

	return deviceLogs, nil
}

func (h *LogQueryHandler) GetWaterLogBySerial(ctx context.Context, query GetWaterLogBySerialQuery) ([]*WaterLog, error) {
	_, err := h.deviceRepository.FindBySerial(ctx, query.DeviceSerial)

	if err != nil {
		return nil, err
	}

	waterLogs, err := h.waterLogRepository.FindBySerial(ctx, query.DeviceSerial)
	if err != nil {
		return nil, err
	}

	return waterLogs, nil
}

func (h *LogQueryHandler) GetWaterLogByFarmId(ctx context.Context, query GetWaterLogByFarmIdQuery) ([]*WaterLog, error) {
	_, err := h.farmRepository.FindById(ctx, query.FarmId)

	if err != nil {
		return nil, err
	}

	fetchedDevices, err := h.deviceRepository.FindByFarmId(ctx, query.FarmId)
	if err != nil {
		return nil, err
	}

	waterLogs := make([]*WaterLog, 0)

	for _, device := range fetchedDevices {
		logs, _ := h.waterLogRepository.FindBySerial(ctx, device.DeviceSerial)
		for _, log := range logs {
			waterLogs = append(waterLogs, log)
		}
	}

	return waterLogs, nil
}
