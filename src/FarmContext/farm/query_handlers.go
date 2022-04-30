package farm


import "context"

type FarmQueryHandler struct {
	farmRepository IRepository
}


func NewFarmQueryHandler(farmRepository IRepository) *FarmQueryHandler {
	return &FarmQueryHandler{farmRepository: farmRepository}
}

func (h *FarmQueryHandler) GetFarmById(ctx context.Context, query GetFarmByIdQuery) (*Farm, error) {
	farm, err := h.farmRepository.FindById(ctx, query.FarmId)
	return farm, err
}


