package water_log



import (
	"farm/src/FarmContext/devices"
	"farm/src/FarmContext/users"
	"gorm.io/gorm"
	"time"

)

type WaterLog struct {
	gorm.Model
	DeviceSerial uint      `gorm:"column:device_serial"`
	Device devices.Device `gorm:"-"`
	//device or farm
	UserId uint      `gorm:"column:user_id"`
	User users.User `gorm:"-"`
	Volume float32 `gorm:"column:water_volume"`
	EntryTime time.Time `gorm:"column:entry_time;" sql:"index:device_time_idx"`

}

func (Farm) TableName() string {
	return "water_log"
}

