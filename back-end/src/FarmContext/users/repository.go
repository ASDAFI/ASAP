package users

import (
	"context"
	"farm/src/infrastructure"
)

type IRepository interface {
	Save(ctx context.Context, user *User) error
	FindByUserName(ctx context.Context, username string) (*User, error)
	FindTokenByUserName(ctx context.Context, username string) (*AuthToken, error)
}

type UserRepository struct {
	dBInfrastructure infrastructure.DBProvider
}

func NewUserRepository(dBInfrastructure infrastructure.DBProvider) *UserRepository {
	return &UserRepository{dBInfrastructure: dBInfrastructure}
}

func (u *UserRepository) Save(ctx context.Context, user *User) error {
	return u.dBInfrastructure.DB.WithContext(ctx).Save(user).Error
}

func (u *UserRepository) FindByUserName(ctx context.Context, username string) (*User, error) {
	user := &User{}
	err := u.dBInfrastructure.DB.WithContext(ctx).Where("username = ?", username).Find(user).Error
	return user, err
}
func (u *UserRepository) FindTokenByUserName(ctx context.Context, username string) (*AuthToken, error) {
	token := &AuthToken{}
	err := u.dBInfrastructure.DB.WithContext(ctx).Where("username = ?", username).Find(token).Error
	return token, err
}
