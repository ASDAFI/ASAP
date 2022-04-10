package service

import (
	ASAP "ASAP/pb/proto"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct {
	userStore  UserStore
	jwtManager *JWTManager
	ASAP.UnimplementedAuthServiceServer
}

func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	return &AuthServer{userStore: userStore, jwtManager: jwtManager}
}
func (server *AuthServer) Login(ctx context.Context, req *ASAP.LoginRequest) (*ASAP.LoginResponse, error)  {
	user,err:=server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal,"cannot find user: %v", err)
	}
	if user==nil || !user.IsCorrectPassword(req.GetPassword()){
		return nil, status.Errorf(codes.NotFound,"incorrect username or password")

	}
	//log.Println(user)
	//log.Println("dfa")
	token,err:=server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal,"cannot generate access token",err)
	}
	res := &ASAP.LoginResponse{AccessToken: token}
	return res, nil
}