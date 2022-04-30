package users

import "context"

type IRepository interface {
	Save(ctx context.Context, user *User) error
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindById(ctx context.Context, userId uint) (*User, error)
}
