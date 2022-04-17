package farm

import "context"

type IRepository interface {
	Save(ctx context.Context, farm *Farm) error
	FindById(ctx context.Context, id uint) (*Farm, error)
}
