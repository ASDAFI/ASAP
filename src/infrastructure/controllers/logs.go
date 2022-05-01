package controllers

import (
	"context"
	"errors"
	"farm/src/FarmContext/devices"
	"farm/src/FarmContext/logs"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"
	pb_log "farm/src/proto/messages/log"
	"github.com/golang/protobuf/ptypes"
	log "github.com/sirupsen/logrus"
)

func (f FarmServer) GetDeviceDataFrameBySerial(ctx context.Context, request *pb_log.GetDeviceDataFrameBySerialRequest) (*pb_log.DeviceDataFrame, error) {
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

	deviceLogRepo := logs.NewDeviceLogRepository(infrastructure.PostgresDBProvider)
	waterLogRepo := logs.NewWaterLogRepository(infrastructure.PostgresDBProvider)

	LogQHandler := logs.NewLogQueryHandler(deviceLogRepo, waterLogRepo, deviceRepo)
	deviceLogQuery := logs.GetDataFrameBySerialQuery{
		DeviceSerial: request.GetDeviceSerial(),
		StartDate:    request.GetStartDate().AsTime(),
		EndDate:      request.GetEndDate().AsTime(),
		Step:         int(request.GetStep()),
	}

	dataFrames, err := LogQHandler.GetDataFrameBySerial(ctx, deviceLogQuery)

	var deviceLogs []*pb_log.DeviceLog
	for _, frame := range dataFrames {
		serverTime, err := ptypes.TimestampProto(frame.ServerTime)
		if err != nil {
			return nil, err
		}
		deviceTime, err := ptypes.TimestampProto(frame.DeviceTime)
		if err != nil {
			return nil, err
		}

		deviceLogs = append(deviceLogs, &pb_log.DeviceLog{
			ServerTime: serverTime,
			DeviceTime: deviceTime,
			Humidity:   uint32(frame.Humidity),
		})

	}

	return &pb_log.DeviceDataFrame{
		DeviceId:     uint32(device.ID),
		DeviceSerial: device.DeviceSerial,
		DeviceLogs:   deviceLogs,
	}, nil

}

func (f FarmServer) GetDeviceDataFrameById(ctx context.Context, request *pb_log.GetDeviceDataFrameByIdRequest) (*pb_log.DeviceDataFrame, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("Receive message to get device by serial ", userId)

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

	deviceLogRepo := logs.NewDeviceLogRepository(infrastructure.PostgresDBProvider)
	waterLogRepo := logs.NewWaterLogRepository(infrastructure.PostgresDBProvider)

	LogQHandler := logs.NewLogQueryHandler(deviceLogRepo, waterLogRepo, deviceRepo)
	deviceLogQuery := logs.GetDataFrameBySerialQuery{
		DeviceSerial: device.DeviceSerial,
		StartDate:    request.GetStartDate().AsTime(),
		EndDate:      request.GetEndDate().AsTime(),
		Step:         int(request.GetStep()),
	}

	dataFrames, err := LogQHandler.GetDataFrameBySerial(ctx, deviceLogQuery)

	var deviceLogs []*pb_log.DeviceLog
	for _, frame := range dataFrames {
		serverTime, err := ptypes.TimestampProto(frame.ServerTime)
		if err != nil {
			return nil, err
		}
		deviceTime, err := ptypes.TimestampProto(frame.DeviceTime)
		if err != nil {
			return nil, err
		}

		deviceLogs = append(deviceLogs, &pb_log.DeviceLog{
			ServerTime: serverTime,
			DeviceTime: deviceTime,
			Humidity:   uint32(frame.Humidity),
		})

	}

	return &pb_log.DeviceDataFrame{
		DeviceId:     uint32(device.ID),
		DeviceSerial: device.DeviceSerial,
		DeviceLogs:   deviceLogs,
	}, nil

}
