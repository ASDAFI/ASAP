
package migration

import (
	"farm/src/FarmContext/devices"
	"farm/src/FarmContext/farm"
	"farm/src/FarmContext/logs"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"

	"gorm.io/gorm"
)


func MigrateDB() error {

	dbProvider := infrastructure.PostgresDBProvider

	err := dbProvider.DB.AutoMigrate(&farm.Farm{}, &users.User{},
		&devices.Device{}, &logs.DeviceLog{}, &logs.WaterLog{}, &users.AuthToken{})
	if err != nil {
		return err
	}
	deleteIndex(dbProvider.DB)
	createUniqueIndex(dbProvider.DB)
	foreignKeyIndex(dbProvider.DB)
	createNormalIndex(dbProvider.DB)
	insertData(dbProvider.DB)
	return nil
}

func insertData(provider *gorm.DB) {

}

func createNormalIndex(provider *gorm.DB) {
}

func deleteIndex(provider *gorm.DB) {
}

func createUniqueIndex(provider *gorm.DB) {
}

func foreignKeyIndex(provider *gorm.DB) {

}

