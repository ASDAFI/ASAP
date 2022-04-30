package farm

import (
	"context"
	"farm/src/infrastructure"
)

type FarmRepository struct {
	dBInfrastructure infrastructure.DBProvider
}

func NewFarmRepository(dBInfrastructure infrastructure.DBProvider) *FarmRepository {
	return &FarmRepository{dBInfrastructure: dBInfrastructure}
}


func (r *FarmRepository) Save(ctx context.Context, farm *Farm) error {
	return r.dBInfrastructure.DB.WithContext(ctx).Save(farm).Error
}


func (r *FarmRepository) FindById(ctx context.Context, farmId uint) (*Farm, error) {
	farm := &Farm{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("id = ?", farmId).Find(farm).Error
	return farm, err

}
