package devices

import (
	"context"
	"errors"
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
	if device.ID != deviceId {
		return nil, errors.New("device not found")
	}
	return device, err

}

func (r *DeviceRepository) FindBySerial(ctx context.Context, deviceSerial string) (*Device, error) {
	device := &Device{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("device_serial = ?", deviceSerial).Find(device).Error
	if device.DeviceSerial != deviceSerial {
		return nil, errors.New("device not found")
	}
	return device, err
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
