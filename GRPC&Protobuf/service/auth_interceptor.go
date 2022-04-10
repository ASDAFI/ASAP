package service

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

type AuthInterceptor struct {
	jwtManager *JWTManager
}

func NewAuthInterceptor(jwtManager *JWTManager) *AuthInterceptor {
	return &AuthInterceptor{jwtManager: jwtManager}
}
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor  {
	return func (ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	log.Println("unary interceptor", info.FullMethod)
	err=interceptor.authorize(ctx,info.FullMethod)
		if err != nil {
			return nil, err
		}
	return handler(ctx,req)
	}
}
func (interceptor *AuthInterceptor) Stream() grpc.StreamServerInterceptor  {
	return func  (srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error{
	log.Println("stream", info.FullMethod)
		err:=interceptor.authorize(ss.Context(),info.FullMethod)
		if err != nil {
			return  err
		}
	return handler(srv,ss)
	}
}
func (interceptor *AuthInterceptor) authorize(ctx context.Context,method string)error {
	//log.Println(method)

	if method=="/GRPC&Protobuf.AuthService/Login"{
		return nil
	}
	md,ok:= metadata.FromIncomingContext(ctx)
	if !ok{
		return status.Errorf(codes.Unauthenticated,"metadata is not provided")
	}

	values:=md["authorization"]
	if len(values)==0{
		return status.Errorf(codes.Unauthenticated,"authentication token is not provided")

	}
	accessToken:=values[0]
	_,err:=interceptor.jwtManager.Verify(accessToken)
	if err != nil {
		return status.Errorf(codes.Unauthenticated,"authentication token is not valid: %v",err)
	}
	return nil

}