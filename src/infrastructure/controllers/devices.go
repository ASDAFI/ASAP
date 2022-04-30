package controllers

import (
	"context"
	"errors"
	"farm/src/FarmContext/devices"
	"farm/src/FarmContext/farm"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"
	pb_device "farm/src/proto/messages/device"
	log "github.com/sirupsen/logrus"
)

func (f FarmServer) CreateDevice(ctx context.Context, request *pb_device.CreateDeviceRequest) (*pb_device.CreateDeviceResponse, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("Receive message to get user: ", userId)

	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQHandler := users.NewUserQueryHandler(userRepo)

	query := users.GetUserByIdQuery{UserId: userId}
	user, err := userQHandler.GetUserById(ctx, query)
	if err != nil {
		return nil, err
	}


	if(user.FarmID != uint(request.GetFarmId())) {
		return nil, errors.New("farm is not owned by this user")
	}

	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)
	farmRepo := farm.NewFarmRepository(infrastructure.PostgresDBProvider)

	deviceCHandler := devices.NewCommandHandler(deviceRepo, farmRepo)
	command := devices.CreateDeviceCommand{DeviceSerial: request.GetDeviceSerial(),
											Phone : request.GetPhone(),
											FarmId: uint(request.GetFarmId()),
											}

	err = deviceCHandler.CreateDevice(ctx, command)

	if err != nil {
		return nil, err
	}

	return &pb_device.CreateDeviceResponse{
		DeviceId: 2000,
	}, nil


}
