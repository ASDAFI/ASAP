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
	"google.golang.org/protobuf/types/known/emptypb"
)

func (f FarmServer) CreateDevice(ctx context.Context, request *pb_device.CreateDeviceRequest) (*pb_device.CreateDeviceResponse, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("Receive message to create device ", userId)

	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQHandler := users.NewUserQueryHandler(userRepo)

	userQuery := users.GetUserByIdQuery{UserId: userId}
	user, err := userQHandler.GetUserById(ctx, userQuery)
	if err != nil {
		return nil, err
	}

	if user.FarmID != uint(request.GetFarmId()) {
		return nil, errors.New("farm is not owned by this user")
	}

	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)
	farmRepo := farm.NewFarmRepository(infrastructure.PostgresDBProvider)

	deviceCHandler := devices.NewCommandHandler(deviceRepo, farmRepo)
	command := devices.CreateDeviceCommand{DeviceSerial: request.GetDeviceSerial(),
		Phone:  request.GetPhone(),
		FarmId: uint(request.GetFarmId()),
	}

	err = deviceCHandler.CreateDevice(ctx, command)

	if err != nil {
		return nil, err
	}

	// Get Device ID
	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)
	devicQuery := devices.GetDeviceBySerialQuery{DeviceSerial: request.GetDeviceSerial()}
	device, err := deviceQHandler.GetDeviceBySerial(ctx, devicQuery)

	return &pb_device.CreateDeviceResponse{
		DeviceId: uint32(device.ID),
	}, nil

}

func (f FarmServer) GetDeviceById(ctx context.Context, request *pb_device.GetDeviceByIdRequest) (*pb_device.Device, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("Receive message to get device by id ", userId)

	// Get Device
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)

	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)
	deviceQuery := devices.GetDeviceByIdQuery{DeviceId: uint(request.GetDeviceId())}

	device, err := deviceQHandler.GetDeviceById(ctx, deviceQuery)
	if err != nil {
		return nil, err
	}

	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQHandler := users.NewUserQueryHandler(userRepo)

	userQuery := users.GetUserByIdQuery{UserId: userId}
	user, err := userQHandler.GetUserById(ctx, userQuery)
	if err != nil {
		return nil, err
	}

	if user.FarmID != device.FarmID {
		return nil, errors.New("device is not owned by this user")
	}

	return &pb_device.Device{
		DeviceId:     uint32(device.ID),
		DeviceSerial: device.DeviceSerial,
		Phone:        device.Phone,
		FarmId:       uint32(device.FarmID),
	}, nil
}

func (f FarmServer) GetDeviceBySerial(ctx context.Context, request *pb_device.GetDeviceBySerialRequest) (*pb_device.Device, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("Receive message to get device by serial ", userId)

	// Get Device
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)

	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)
	deviceQuery := devices.GetDeviceBySerialQuery{DeviceSerial: request.GetDeviceSerial()}

	device, err := deviceQHandler.GetDeviceBySerial(ctx, deviceQuery)
	if err != nil {
		return nil, err
	}

	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQHandler := users.NewUserQueryHandler(userRepo)

	userQuery := users.GetUserByIdQuery{UserId: userId}
	user, err := userQHandler.GetUserById(ctx, userQuery)
	if err != nil {
		return nil, err
	}

	if user.FarmID != device.FarmID {
		return nil, errors.New("device is not owned by this user")
	}

	return &pb_device.Device{
		DeviceId:     uint32(device.ID),
		DeviceSerial: device.DeviceSerial,
		Phone:        device.Phone,
		FarmId:       uint32(device.FarmID),
	}, nil
}

func (f FarmServer) GetDevices(ctx context.Context, empty *emptypb.Empty) (*pb_device.Devices, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("Receive message to get devices ", userId)

	// Get User
	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQHandler := users.NewUserQueryHandler(userRepo)

	userQuery := users.GetUserByIdQuery{UserId: userId}
	user, err := userQHandler.GetUserById(ctx, userQuery)
	if err != nil {
		return nil, err
	}

	// Get Device
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)

	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)
	deviceQuery := devices.GetDeviceByFarmIdQuery{FarmId: user.FarmID}

	fetchedDevices, err := deviceQHandler.GetDeviceByFarmId(ctx, deviceQuery)
	if err != nil {
		return nil, err
	}

	var deviceList []*pb_device.Device
	for _, device := range fetchedDevices {
		deviceList = append(deviceList, &pb_device.Device{
			DeviceId:     uint32(device.ID),
			DeviceSerial: device.DeviceSerial,
			Phone:        device.Phone,
			FarmId:       uint32(device.FarmID),
		})
	}

	return &pb_device.Devices{
		Devices: deviceList,
	}, nil
}
