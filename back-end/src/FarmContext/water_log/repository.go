package water_log

import (
	"context"
	"farm/src/infrastructure"
)

type IRepository interface {
	Save(ctx context.Context, log *WaterLog) error
	FindById(ctx context.Context, id uint) (*WaterLog, error)
	Updates(ctx context.Context,log *WaterLog) error
}
type WaterLogRepository struct {
	dBInfrastructure infrastructure.DBProvider
}

func NewWaterLogRepository(dBInfrastructure infrastructure.DBProvider) *WaterLogRepository {
	return &WaterLogRepository{dBInfrastructure: dBInfrastructure}
}

func (waterLog *WaterLogRepository) Save(ctx context.Context, log *WaterLog) error {
	return waterLog.dBInfrastructure.DB.WithContext(ctx).Save(log).Error
}

func (waterLog *WaterLogRepository) FindById(ctx context.Context, id uint) (*WaterLog, error) {
	log := &WaterLog{}
	err := waterLog.dBInfrastructure.DB.WithContext(ctx).Where("ID = ?", id).Find(log).Error
	return log, err
}
func (waterLog *WaterLogRepository) Updates(ctx context.Context, log *WaterLog) error {
	return waterLog.dBInfrastructure.DB.WithContext(ctx).Updates(log).Error

}