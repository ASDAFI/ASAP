package infrastructure

import (
	"fmt"
	"gorm.io/driver/postgres"
	"time"

	"farm/configs"
	log "github.com/sirupsen/logrus"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)



type DBProvider struct {
	config configs.DatabaseConfiguration

	DB *gorm.DB
}

var PostgresDBProvider DBProvider


func CreateDBProvider(config configs.DatabaseConfiguration) (DBProvider, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DB)
	fmt.Printf(connectionString)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	provider := &DBProvider{
		config: config,
		DB:     db,
	}
	if err != nil {
		log.Info("Error in Create db: ", err)
		return *provider, err
	}

	db1, err := db.DB()
	if err != nil {
		log.Info("Error in Create db: ", err)
		return *provider, err
	}
	db1.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetime))
	db1.SetMaxIdleConns(config.MaxIdleConns)
	db1.SetMaxOpenConns(config.MaxOpenConns)
	//db.LogMode(true)
	//logger := log.New()
	//logger.SetLevel(log.DebugLevel)
	//db.SetLogger(logger)

	log.Info("Create Db Done.")
	return *provider, nil
}

