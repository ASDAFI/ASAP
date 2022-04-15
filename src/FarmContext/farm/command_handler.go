package farm

import (
	"context"
)

type CommandHandler struct {
	repository IRepository
}

func NewCommandHandler(farmRepository IRepository) *CommandHandler {
	return &CommandHandler{farmRepository}
}

func (h CommandHandler) Create(ctx context.Context, farmName string) (*Farm, error) {
	farm := &Farm{
		FarmName: farmName,
	}
	err := h.repository.Save(ctx, farm)
	return farm, err
}
