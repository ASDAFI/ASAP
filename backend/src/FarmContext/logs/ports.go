package logs

import (
	"context"
	"time"
)

type IDeviceLogRepository interface {
	Save(ctx context.Context, deviceLog *DeviceLog) error
	GetDataFrameByDeviceId(ctx context.Context, deviceId uint, startDate time.Time, endDate time.Time) ([]*DeviceLog, error)
	GetDataFrameBySerial(ctx context.Context, deviceSerial string, startDate time.Time, endDate time.Time) ([]*DeviceLog, error)
	GetAllDataFramesByDeviceId(ctx context.Context, deviceId uint) ([]*DeviceLog, error)
	GetAllDataFramesBySerial(ctx context.Context, deviceSerial string) ([]*DeviceLog, error)
}

type IWaterLogRepository interface {
	Save(ctx context.Context, waterLog *WaterLog) error
	FindBySerial(ctx context.Context, deviceSerial string) ([]*WaterLog, error)
	FindById(ctx context.Context, Id uint) (*WaterLog, error)
}
