package client

import (
	ASAP "ASAP/pb/proto"
	"context"
	"google.golang.org/grpc"
	"time"
)

type AuthClient struct {
	service  ASAP.AuthServiceClient
	username string
	password string
}

func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := ASAP.NewAuthServiceClient(cc)
	return &AuthClient{service: service, username: username, password: password}
}

func (client *AuthClient) Login() (string,error) {
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()
	req:= &ASAP.LoginRequest{
		Username: client.username,
		Password: client.password,
	}
	res,err:=client.service.Login(ctx,req)
	if err != nil {
		return "", err
	}
	return res.GetAccessToken(),nil
}
