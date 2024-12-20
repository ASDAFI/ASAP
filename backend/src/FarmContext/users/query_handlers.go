package users

import "context"

type UserQueryHandler struct {
	userRepository IRepository
}


func NewUserQueryHandler(userRepository IRepository) *UserQueryHandler {
	return &UserQueryHandler{
		userRepository: userRepository,
	}
}

func (h *UserQueryHandler) GetUserByUsername(ctx context.Context, query GetUserByUsernameQuery) (*User, error) {
	user, err := h.userRepository.FindByUsername(ctx, query.Username)
	return user, err
}

func (h *UserQueryHandler) GetUserById(ctx context.Context, query GetUserByIdQuery) (*User, error) {
	user, err := h.userRepository.FindById(ctx, query.UserId)
	return user, err
}

