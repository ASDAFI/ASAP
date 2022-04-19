package water_log

import (
"context"
	"farm/src/FarmContext/devices"
	"farm/src/FarmContext/users"
)

type CommandHandler struct {
	repository IRepository
}

func NewCommandHandler(farmRepository IRepository) *CommandHandler {
	return &CommandHandler{farmRepository}
}

type CreateWaterLog struct {
	DeviceSerial uint
	UserId uint
	Volume float32
	EntryTime time.Time
}
func (h CommandHandler) Create(ctx context.Context, waterLog CreateWaterLog) (*WaterLog, error) {
	newWaterLog := &WaterLog{
		DeviceSerial: waterLog.DeviceSerial,
		Device:       devices.Device{},
		UserId:       waterLog.UserId,
		User:         users.User{},
		Volume:       waterLog.Volume,
		EntryTime:    waterLog.EntryTime,
	}
	err := h.repository.Save(ctx, newWaterLog)
	return newWaterLog, err
}

