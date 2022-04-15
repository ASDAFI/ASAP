package users

import (
	"context"
	"farm/src/FarmContext/farm"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type CommandHandler struct {
	farmRepo       farm.IRepository
	UserRepository IRepository
}

func NewCommandHandler(farmRepo farm.IRepository, userRepository IRepository) *CommandHandler {
	return &CommandHandler{farmRepo: farmRepo, UserRepository: userRepository}
}

type CreateUserCommand struct {
	Username string
	Password string
	FarmID   uint
}

func (h CommandHandler) Create(ctx context.Context, command CreateUserCommand) (*User, error) {
	// todo : use HashService
	pass, err := bcrypt.GenerateFromPassword([]byte(command.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Info("Password Encryption  failed")
		return nil, err
	}

	user, err := NewUser(CreateParameters{
		Username: command.Username,
		Password: string(pass),
	})
	if err != nil {
		return nil, err
	}

	farmModel, err := h.farmRepo.FindById(ctx, command.FarmID)
	if err != nil {
		return user, err
	}
	user.AssignFarm(*farmModel)
	err = h.UserRepository.Save(ctx, user)
	return user, err

}
