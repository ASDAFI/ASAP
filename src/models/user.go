package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName  string `gorm:"column:first_name"`
	LastName   string `gorm:"column:last_name"`
	Username   string `gorm:"column:username;unique_index;not null;"`
	Password   string `gorm:"column:password"`
	NationalID string `gorm:"unique_index;not null;"`
	Email      string `gorm:"column:email"`
	Phone      string `gorm:"column:phone"`
	FarmID     uint   `gorm:"column:company_id"`
	Farm       Farm
}

type AuthToken struct {
	Username string `gorm:"column:username"`
	*jwt.StandardClaims
}

func (User) TableName() string {
	return "auth_user"
}

func (AuthToken) TableName() string {
	return "auth_token"
}

func StandardClaims(expiresAt int64, id string) *jwt.StandardClaims {
	return &jwt.StandardClaims{
		ExpiresAt: expiresAt,
		Id:        id,
	}
}
