package devices

import (
	"context"
	"gorm.io/gorm"
)

type WaterLogRepository struct {
	db *gorm.DB
}

func NewWaterLogRepository(db *gorm.DB) *WaterLogRepository {
	return &WaterLogRepository{db: db}
}

func (r *WaterLogRepository) Save(ctx context.Context, log *WaterLog) error {
	return r.db.Save(log).Error
}
func (r *WaterLogRepository) Updates(ctx context.Context, log *WaterLog) error {
	return r.db.Updates(log).Error
}

type DeviceRepository struct {
	db *gorm.DB
}

func NewDeviceRepository(db *gorm.DB) *DeviceRepository {
	return &DeviceRepository{db: db}
}

func (r *DeviceRepository) FindById(ctx context.Context, deviceId uint) (*Device, error) {
	log := &Device{}
	err := r.db.WithContext(ctx).Where("ID = ?", deviceId).Find(log).Error
	return log, err
}

func (r DeviceRepository) FindBySerial(ctx context.Context, deviceSerial string) (*Device, error) {
	log := &Device{}
	err := r.db.WithContext(ctx).Where("device_serial = ?", deviceSerial).Find(log).Error
	return log, err
}
