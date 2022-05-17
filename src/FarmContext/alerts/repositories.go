package alerts

import (
	"context"
	"farm/src/infrastructure"
)

type AlertRepository struct {
	dBInfrastructure infrastructure.DBProvider
}

func NewAlertRepository(dBInfrastructure infrastructure.DBProvider) *IRepository {
	return &AlertRepository{dBInfrastructure: dBInfrastructure}
}

func (r *AlertRepository) Save(ctx context.Context, alert *Alert) error {
	return r.dBInfrastructure.DB.WithContext(ctx).Save(alert).Error
}
