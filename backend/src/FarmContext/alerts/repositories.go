package alerts

import (
	"context"
	"farm/src/infrastructure"
)

type AlertRepository struct {
	dBInfrastructure infrastructure.DBProvider
}

func NewAlertRepository(dBInfrastructure infrastructure.DBProvider) *AlertRepository {
	return &AlertRepository{dBInfrastructure: dBInfrastructure}
}

func (r *AlertRepository) Save(ctx context.Context, alert *Alert) error {
	return r.dBInfrastructure.DB.WithContext(ctx).Save(alert).Error
}

func (r *AlertRepository) FindAll(ctx context.Context) ([]*Alert, error) {
	var alerts []*Alert
	err := r.dBInfrastructure.DB.WithContext(ctx).Find(&alerts).Error
	return alerts, err
}

func (r *AlertRepository) FindLastByDeviceSerial(ctx context.Context, deviceSerial string) (*Alert, error) {
	var alert Alert
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("device_serial = ?", deviceSerial).Last(&alert).Error
	return &alert, err
}

func (r *AlertRepository) FindByDeviceSerial(ctx context.Context, deviceSerial string) ([]*Alert, error) {
	var alerts []*Alert
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("device_serial = ?", deviceSerial).Find(&alerts).Error
	return alerts, err
}

func (r *AlertRepository) FindRecentByDeviceSerial(ctx context.Context, deviceSerial string) ([]*Alert, error) {
	var alerts []*Alert
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("device_serial = ?", deviceSerial).Where("(now()::timestamp - created_at::timestamp )< '10 secs'::interval").Find(&alerts).Error
	return alerts, err
}
