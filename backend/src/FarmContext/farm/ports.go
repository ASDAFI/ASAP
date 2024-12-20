package farm

import "context"

type IRepository interface {
	Save(ctx context.Context, farm *Farm) error
	FindById(ctx context.Context, farmId uint) (*Farm, error)
}

