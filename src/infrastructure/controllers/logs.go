package controllers

import (
	"context"
	"errors"
	"farm/src/FarmContext/devices"
	"farm/src/FarmContext/farm"
	"farm/src/FarmContext/logs"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"
	pb_log "farm/src/proto/messages/log"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
)

func (f FarmServer) GetDeviceDataFrameBySerial(ctx context.Context, request *pb_log.GetDeviceDataFrameBySerialRequest) (*pb_log.DeviceDataFrame, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("GetDeviceDataFrameBySerial -- userId: ", userId)

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
	farmRepo := farm.NewFarmRepository(infrastructure.PostgresDBProvider)

	logQHandler := logs.NewLogQueryHandler(deviceLogRepo, waterLogRepo, deviceRepo, farmRepo)
	deviceLogQuery := logs.GetDataFrameBySerialQuery{
		DeviceSerial: request.GetDeviceSerial(),
		StartDate:    request.GetStartDate().AsTime(),
		EndDate:      request.GetEndDate().AsTime(),
		Step:         int(request.GetStep()),
	}

	dataFrames, err := logQHandler.GetDataFrameBySerial(ctx, deviceLogQuery)

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
	log.Info("GetDeviceDataFrameById  -- userId: ", userId)

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
	farmRepo := farm.NewFarmRepository(infrastructure.PostgresDBProvider)

	logQHandler := logs.NewLogQueryHandler(deviceLogRepo, waterLogRepo, deviceRepo, farmRepo)
	deviceLogQuery := logs.GetDataFrameBySerialQuery{
		DeviceSerial: device.DeviceSerial,
		StartDate:    request.GetStartDate().AsTime(),
		EndDate:      request.GetEndDate().AsTime(),
		Step:         int(request.GetStep()),
	}

	dataFrames, err := logQHandler.GetDataFrameBySerial(ctx, deviceLogQuery)

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

func (f FarmServer) GetAllDeviceDataFramesBySerial(ctx context.Context, request *pb_log.GetDeviceDataFrameBySerialRequest) (*pb_log.DeviceDataFrame, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("GetAllDeviceDataFramesBySerial  -- userId: ", userId)

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
	farmRepo := farm.NewFarmRepository(infrastructure.PostgresDBProvider)

	logQHandler := logs.NewLogQueryHandler(deviceLogRepo, waterLogRepo, deviceRepo, farmRepo)
	deviceLogQuery := logs.GetAllDataFramesBySerialQuery{
		DeviceSerial: request.GetDeviceSerial(),
	}

	dataFrames, err := logQHandler.GetAllDataFramesBySerial(ctx, deviceLogQuery)

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

func (f FarmServer) GetAllDeviceDataFramesById(ctx context.Context, request *pb_log.GetDeviceDataFrameByIdRequest) (*pb_log.DeviceDataFrame, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("GetAllDeviceDataFramesById  -- userId: ", userId)

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
	farmRepo := farm.NewFarmRepository(infrastructure.PostgresDBProvider)

	logQHandler := logs.NewLogQueryHandler(deviceLogRepo, waterLogRepo, deviceRepo, farmRepo)
	deviceLogQuery := logs.GetAllDataFramesBySerialQuery{
		DeviceSerial: device.DeviceSerial,
	}

	dataFrames, err := logQHandler.GetAllDataFramesBySerial(ctx, deviceLogQuery)

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

func (f FarmServer) CreateWaterLogBySerial(ctx context.Context, request *pb_log.CreateWaterLogBySerialRequest) (*empty.Empty, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("CreateWaterLogBySerial  -- userId: ", userId)

	// Get Device
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)

	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)
	deviceQuery := devices.GetDeviceBySerialQuery{DeviceSerial: request.GetDeviceSerial()}

	_, err := deviceQHandler.GetDeviceBySerial(ctx, deviceQuery)
	if err != nil {
		return nil, errors.New("device not found")
	}

	waterLogRepo := logs.NewWaterLogRepository(infrastructure.PostgresDBProvider)
	deviceLogRepo := logs.NewDeviceLogRepository(infrastructure.PostgresDBProvider)

	logCHandler := logs.NewCommandHandler(deviceRepo, deviceLogRepo, waterLogRepo)
	waterLogCommand := logs.CreateWaterLogCommand{
		DeviceSerial: request.GetDeviceSerial(),
		UserId:       userId,
		Volume:       uint(request.GetVolume()),
		EntryTime:    request.GetEntryTime().AsTime(),
	}

	err = logCHandler.CreateWaterLog(ctx, waterLogCommand)

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil

}

func (f FarmServer) CreateWaterLogByDeviceId(ctx context.Context, request *pb_log.CreateWaterLogByDeviceIdRequest) (*empty.Empty, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("CreateWaterLogByDeviceId -- userId: ", userId)

	// Get Device
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)

	deviceQHandler := devices.NewDeviceQueryHandler(deviceRepo)
	deviceQuery := devices.GetDeviceByIdQuery{DeviceId: uint(request.GetDeviceId())}

	device, err := deviceQHandler.GetDeviceById(ctx, deviceQuery)
	if err != nil || device == nil {
		return nil, errors.New("device not found")
	}

	waterLogRepo := logs.NewWaterLogRepository(infrastructure.PostgresDBProvider)
	deviceLogRepo := logs.NewDeviceLogRepository(infrastructure.PostgresDBProvider)

	logCHandler := logs.NewCommandHandler(deviceRepo, deviceLogRepo, waterLogRepo)
	waterLogCommand := logs.CreateWaterLogCommand{
		DeviceSerial: device.DeviceSerial,
		UserId:       userId,
		Volume:       uint(request.GetVolume()),
		EntryTime:    request.GetEntryTime().AsTime(),
	}

	err = logCHandler.CreateWaterLog(ctx, waterLogCommand)

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil

}

func (f FarmServer) GetWaterLogs(ctx context.Context, empty *empty.Empty) (*pb_log.GetWaterLogsResponse, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("GetWaterLogs -- userId: ", userId)

	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQhandler := users.NewUserQueryHandler(userRepo)
	userQuery := users.GetUserByIdQuery{UserId: userId}
	user, err := userQhandler.GetUserById(ctx, userQuery)

	if err != nil {
		return nil, err
	}

	waterLogRepo := logs.NewWaterLogRepository(infrastructure.PostgresDBProvider)
	deviceRepo := devices.NewDeviceRepository(infrastructure.PostgresDBProvider)
	deviceLogRepo := logs.NewDeviceLogRepository(infrastructure.PostgresDBProvider)
	farmRepo := farm.NewFarmRepository(infrastructure.PostgresDBProvider)

	logQHandler := logs.NewLogQueryHandler(deviceLogRepo, waterLogRepo, deviceRepo, farmRepo)

	logQuery := logs.GetWaterLogByFarmIdQuery{FarmId: user.FarmID}

	waterLogs, err := logQHandler.GetWaterLogByFarmId(ctx, logQuery)

	if err != nil {
		return nil, err
	}

	waterLogsResponse := make([]*pb_log.WaterLog, len(waterLogs))

	for i, waterLog := range waterLogs {
		userQuery = users.GetUserByIdQuery{UserId: waterLog.UserId}
		user, err := userQhandler.GetUserById(ctx, userQuery)

		if err != nil {
			return nil, err
		}

		waterLogsResponse[i] = &pb_log.WaterLog{
			Id:           uint32(waterLog.ID),
			DeviceSerial: waterLog.DeviceSerial,
			Volume:       uint32(waterLog.Volume),
			EntryTime:    ptypes.TimestampNow(),
			Username:     user.Username,
		}
	}

	return &pb_log.GetWaterLogsResponse{
		FarmId:    uint32(user.FarmID),
		WaterLogs: waterLogsResponse,
	}, nil
}
