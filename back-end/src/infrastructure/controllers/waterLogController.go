package controllers

import (
	"context"
	"farm/src/FarmContext/devices"
	"farm/src/infrastructure"
	"farm/src/proto/messages/water_log"
)

type WaterLogServer struct {

}
func (w *WaterLogServer) CreateWaterLog(ctx context.Context, request *water_log.CreateWaterLogRequest) (*water_log.CreateWaterLogResponse, error){
	userId := ctx.Value("userId").(uint)
	commandHandler:= devices.NewDeviceCommandHandler(devices.NewWaterLogRepository(infrastructure.PostgresDBProvider.DB),devices.NewDeviceRepository(infrastructure.PostgresDBProvider.DB))
	err := commandHandler.AddWaterHandler(ctx , devices.AddWaterCommand{
		DeviceId: uint(request.GetDeviceId()),
		UserId:   userId,
		Volume:   request.GetVolume(),
		Time:     request.GetEntryTime().AsTime(),
	})
	return nil, err

}