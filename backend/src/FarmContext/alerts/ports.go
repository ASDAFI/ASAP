package alerts

import "context"

type IRepository interface {
	Save(ctx context.Context, alert *Alert) error
	FindAll(ctx context.Context) ([]*Alert, error)
	FindLastByDeviceSerial(ctx context.Context, deviceSerial string) (*Alert, error)
	FindByDeviceSerial(ctx context.Context, deviceSerial string) ([]*Alert, error)
	FindRecentByDeviceSerial(ctx context.Context, deviceSerial string) ([]*Alert, error)
	//GetAllAlertsForDevice(ctx context.Context, serial string) ([]*Alert, error)
}
