package models

import (
	"fmt"
	"sync"
	"time"

	"farm/configs"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var PostgresDBProvider DBProvider

type DBProvider struct {
	config configs.DatabaseConfiguration

	mutex sync.Mutex
	DB    *gorm.DB
}

func CreateDBProvider(config configs.DatabaseConfiguration) (DBProvider, error) {
	provider := &DBProvider{
		config: config,
	}

	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s", config.Host, config.Port, config.User, config.Password, config.DB))

	if err != nil {
		log.Info("Error in Create PostgresDBProvider: ", err)
		return *provider, err
	}
	db.DB().SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetime))
	db.DB().SetMaxIdleConns(config.MaxIdleConns)
	db.DB().SetMaxOpenConns(config.MaxOpenConns)

	provider.DB = db
	return *provider, nil
}

func (provider *DBProvider) MigrateDB() error {
	provider.DB.LogMode(true)
	err := provider.DB.AutoMigrate(&Farm{}, &User{},
		&Device{}, &DeviceLog{}, &AuthToken{}).Error
	if err != nil {
		return err
	}
	deleteIndex(provider)
	createUniqueIndex(provider)
	foreignKeyIndex(provider)
	createNormalIndex(provider)
	insertData(provider)
	return nil
}
func insertData(provider *DBProvider) {

}

func createNormalIndex(provider *DBProvider) {
}

func deleteIndex(provider *DBProvider) {
}

func createUniqueIndex(provider *DBProvider) {
}

func foreignKeyIndex(provider *DBProvider) {

}
