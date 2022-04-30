package logs

import (
	"context"
	"farm/src/FarmContext/devices"
)

type LogQueryHandler struct {
	deviceLogRepository IDeviceLogRepository
	waterLogRepository IWaterLogRepository
	deviceRepository   devices.IRepository
}


func NewLogQueryHandler(deviceLogRepository IDeviceLogRepository, waterLogRepository IWaterLogRepository,
					    deviceRepository devices.IRepository ) *LogQueryHandler {
	return &LogQueryHandler{
		deviceLogRepository: deviceLogRepository,
		waterLogRepository: waterLogRepository,
		deviceRepository: deviceRepository,
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
	for i :=0 ; i < len(dataFrame); i = i + query.Step {
		dataFrame = append(dataFrame, deviceLogs[i])
	}

	return dataFrame, nil
}

func (h *LogQueryHandler) GetDataFrameBySerial(ctx context.Context, query *GetDataFrameBySerialQuery) ([]*DeviceLog, error) {
	_, err := h.deviceRepository.FindBySerial(ctx, query.DeviceSerial)

	if err != nil {
		return nil, err
	}

	deviceLogs, err := h.deviceLogRepository.GetDataFrameBySerial(ctx, query.DeviceSerial, query.StartDate, query.EndDate)
	if err != nil {
		return nil, err
	}

	dataFrame := make([]*DeviceLog, 0)
	for i :=0 ; i < len(dataFrame); i = i + query.Step {
		dataFrame = append(dataFrame, deviceLogs[i])
	}

	return dataFrame, nil
}

func (h *LogQueryHandler) GetWaterLogByDeviceId(ctx context.Context, query *GetWaterLogByDeviceIdQuery) ([]*WaterLog, error) {
	_, err := h.deviceRepository.FindById(ctx, query.DeviceId)

	if err != nil {
		return nil, err
	}

	waterLogs, err := h.waterLogRepository.FindByDeviceId(ctx, query.DeviceId)
	if err != nil {
		return nil, err
	}

	return waterLogs, nil
}

func (h *LogQueryHandler) GetWaterLogBySerial(ctx context.Context, query *GetWaterLogBySerialQuery) ([]*WaterLog, error) {
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