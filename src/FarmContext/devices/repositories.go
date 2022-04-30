package devices

import (
	"context"
	"farm/src/infrastructure"
)

type DeviceRepository struct {
	dBInfrastructure infrastructure.DBProvider
}

func NewDeviceRepository(dBInfrastructure infrastructure.DBProvider) *DeviceRepository {
	return &DeviceRepository{dBInfrastructure: dBInfrastructure}
}


func (r *DeviceRepository) Save(ctx context.Context, device *Device) error {
	return r.dBInfrastructure.DB.WithContext(ctx).Save(device).Error
}


func (r *DeviceRepository) FindById(ctx context.Context, deviceId uint) (*Device, error) {
	device := &Device{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("id = ?", deviceId).Find(device).Error
	return device, err

}


func (r *DeviceRepository) FindBySerial(ctx context.Context, deviceSerial string) (*Device, error) {
	log := &Device{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("device_serial = ?", deviceSerial).Find(log).Error
	return log, err
}

func (r *DeviceRepository) GetAll(ctx context.Context) ([]*Device, error) {
	devices := make([]*Device, 0)
	err := r.dBInfrastructure.DB.WithContext(ctx).Find(&devices).Error
	return devices, err
}

func (r *DeviceRepository) FindByFarmId(ctx context.Context, farmId uint) ([]*Device, error) {
	devices := make([]*Device, 0)
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("farm_id = ?", farmId).Find(&devices).Error
	return devices, err
}