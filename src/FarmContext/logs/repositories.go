package logs

import (
	"context"
	"farm/src/infrastructure"
	"time"
)

type DeviceLogRepository struct {
	dBInfrastructure infrastructure.DBProvider
}

type WaterLogRepository struct {
	dBInfrastructure infrastructure.DBProvider
}

func NewDeviceLogRepository(dBInfrastructure infrastructure.DBProvider) *DeviceLogRepository {
	return &DeviceLogRepository{dBInfrastructure: dBInfrastructure}
}

func NewWaterLogRepository(dBInfrastructure infrastructure.DBProvider) *WaterLogRepository {
	return &WaterLogRepository{dBInfrastructure: dBInfrastructure}
}

func (r *DeviceLogRepository) Save(ctx context.Context, deviceLog *DeviceLog) error {
	return r.dBInfrastructure.DB.WithContext(ctx).Save(deviceLog).Error
}

func (r *DeviceLogRepository) GetDataFrameByDeviceId(ctx context.Context, deviceId uint,
	                                       startTime time.Time, endTime time.Time) ([]*DeviceLog, error) {
	logs := []*DeviceLog{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("device_id = ? AND datetime BETWEEN ? AND ? ", deviceId,
		                                            startTime, endTime).Order("datetime desc").Find(&logs).Error
	return logs, err
}

func (r *DeviceLogRepository) GetDataFrameBySerial(ctx context.Context, deviceSerial string,
	startTime time.Time, endTime time.Time) ([]*DeviceLog, error) {
	logs := []*DeviceLog{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("device_serial = ? AND datetime BETWEEN ? AND ? ", deviceSerial,
		startTime, endTime).Order("datetime desc").Find(&logs).Error
	return logs, err
}

func (r *WaterLogRepository) Save(ctx context.Context, waterLog *WaterLog) error {
	return r.dBInfrastructure.DB.WithContext(ctx).Save(waterLog).Error
}

func (r *WaterLogRepository) FindById(ctx context.Context, waterLogId uint) (*WaterLog, error) {
	waterLog := &WaterLog{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("id = ?", waterLogId).Find(waterLog).Error
	return waterLog, err

}


func (r *WaterLogRepository) FindByDeviceId(ctx context.Context, deviceId uint) ([]*WaterLog, error) {
	waterLogs := []*WaterLog{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("device_id = ?", deviceId).Find(&waterLogs).Error
	return waterLogs, err
}

func (r *WaterLogRepository) FindBySerial(ctx context.Context, deviceSerial string) ([]*WaterLog, error) {
	waterLogs := []*WaterLog{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("device_serial = ?", deviceSerial).Find(&waterLogs).Error
	return waterLogs, err
}