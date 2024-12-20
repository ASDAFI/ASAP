package devices


import "context"

type IRepository interface {
	Save(ctx context.Context, user *Device) error
	FindById(ctx context.Context, deviceId uint) (*Device, error)
	FindBySerial(ctx context.Context, serial string) (*Device, error)
	GetAll(ctx context.Context) ([]*Device, error)
	FindByFarmId(ctx context.Context, farmId uint) ([]*Device, error)
}
