package controllers

import (
	"context"
	"farm/src/FarmContext/users"
	"farm/src/infrastructure"
	pb_user "farm/src/proto/messages/user"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (f FarmServer) Login(ctx context.Context, request *pb_user.LoginRequest) (*pb_user.LoginResponse, error) {
	log.Info("Receive message to login: ", request.Username)
	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	authHandler := users.NewAuthHandler(userRepo)
	loginErr, token := authHandler.Login(ctx, request.GetUsername(), request.GetPassword())
	return &pb_user.LoginResponse{
		Token: token,
	}, loginErr
}

func (f FarmServer) GetUser(ctx context.Context, empty *emptypb.Empty) (*pb_user.User, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("Receive message to get user: ", userId)

	userRepo := users.NewUserRepository(infrastructure.PostgresDBProvider)
	userQHandler := users.NewUserQueryHandler(userRepo)

	query := users.GetUserByIdQuery{UserId: userId}
	user, err := userQHandler.GetUserById(ctx, query)
	if err != nil {
		return nil, err
	}
	return &pb_user.User{
		Username:  user.Username,
		Email:     user.Email,
		Phone:     user.Phone,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil

}
