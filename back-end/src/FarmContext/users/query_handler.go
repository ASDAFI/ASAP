package users

import (
	"context"
)

type UserQueryHandler struct {
	userRepository IRepository
}

func (h *UserQueryHandler) GetUserByName(ctx context.Context, username string) (*User, error) {
	user, err := h.userRepository.FindByUserName(ctx, username)
	return user, err
}
