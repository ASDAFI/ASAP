package users

import (
	"context"
	"farm/src/infrastructure"
)

type UserRepository struct {
	dBInfrastructure infrastructure.DBProvider
}

func NewUserRepository(dBInfrastructure infrastructure.DBProvider) *UserRepository {
	return &UserRepository{dBInfrastructure: dBInfrastructure}
}

func (r *UserRepository) Save(ctx context.Context, user *User) error {
	return r.dBInfrastructure.DB.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*User, error) {
	user := &User{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("username = ?", username).Find(user).Error
	return user, err
}

func (r *UserRepository) FindById(ctx context.Context, userId uint) (*User, error) {
	user := &User{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("id = ?", userId).Find(user).Error
	return user, err

}

func (r *UserRepository) FindTokenByUsername(ctx context.Context, username string) (*AuthToken, error) {
	token := &AuthToken{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("username = ?", username).Find(token).Error
	return token, err
}
