package alerts

import "context"

type IRepository interface {
	Save(ctx context.Context, user *Alert) error
	GetAllAlertsForDevice(ctx context.Context, serial string) ([]*Alert, error)
}
