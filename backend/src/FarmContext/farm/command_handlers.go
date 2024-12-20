package farm

import (
	"context"
)

type CommandHandler struct {
	farmRepo IRepository
}

func NewCommandHandler(farmRepo IRepository) *CommandHandler {
	return &CommandHandler{farmRepo: farmRepo}
}


func (h CommandHandler) CreateFarm(ctx context.Context, command CreateFarmCommand) (*Farm, error) {

	farm, err := NewFarm(CreateFarmParameters{
		FarmName: command.FarmName,
		Production: command.Production,
	})
	if err != nil {
		return nil, err
	}


	err = h.farmRepo.Save(ctx, farm)
	return farm, err

}
