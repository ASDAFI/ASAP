package service

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username       string
	HashedPassword string
}

func NewUser(username string, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}
	user := &User{
		Username:       username,
		HashedPassword: string(hashedPassword),
	}
	return user, nil
}
func (user *User) IsCorrectPassword(password string)bool  {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err==nil
}
func (user *User) Clone()  *User {
return &User{
	Username:       user.Username,
	HashedPassword: user.HashedPassword,
}
}