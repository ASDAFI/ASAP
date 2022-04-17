package controllers

import (
	"context"
	"farm/src/FarmContext/users"
	pb_user "farm/src/proto/messages/user"
	pb_device "farm/src/proto/messages/devices"
	log "github.com/sirupsen/logrus"
)

type FarmServer struct {
}

func (f FarmServer) Login(ctx context.Context, request *pb_user.LoginRequest) (*pb_user.LoginResponse, error) {
	log.Info("Receive message to login: ", request.Username)
	authHandler := users.AuthHandler{}
	loginErr, token := authHandler.Login(ctx, request.GetUsername(), request.GetPassword())
	return &pb_user.LoginResponse{
		Token: token,
	}, loginErr
}

func (f FarmServer) CreateDevice(ctx context.Context, request *pb_device.CreateDeviceRequest) (*pb_device.CreateDeviceResponse, error) {
	log.Info("Receive message to create device")

	return nil, nil
}
