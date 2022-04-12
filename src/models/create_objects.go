package models

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func CreateFarm(farmname string) (uint, error) {
	farm := &Farm{
		FarmName: farmname,
	}
	if err := PostgresDBProvider.DB.Create(farm).Error; err != nil {
		return 0, err
	}
	return farm.ID, nil
}

func CreateUser(username string, password string, farmID uint) error {
	user := &User{}
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Info("Password Encryption  failed")
		return err
	}
	user.Password = string(pass)
	user.Username = username
	user.FarmID = uint(farmID)
	createdUser := PostgresDBProvider.DB.Create(user)

	if createdUser.Error != nil {
		log.Info(createdUser.Error)
		return createdUser.Error
	}
	return nil
}
