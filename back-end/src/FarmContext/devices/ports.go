package devices

import "context"

type IDeviceRepository interface {
	FindById(ctx context.Context, deviceId uint) (*Device,error)
	FindBySerial(ctx context.Context, deviceSerial string) (*Device,error)
}

type IWaterLogRepository interface {
	Save(ctx context.Context, log *WaterLog) error
	Updates(ctx context.Context, log *WaterLog) error
}
