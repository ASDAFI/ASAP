package client

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"time"
)

type AuthInterceptor struct {
	authClient *AuthClient
	authMethods map[string]bool
	accessToken string
}

func (interceptor *AuthInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {
	err:= interceptor.refreshToken()
	if err != nil {
		return err
	}
	go func() {wait:=refreshDuration
	for{
		time.Sleep(wait)
		err:=interceptor.refreshToken()
		if err != nil {
			wait=time.Second
		}else{
			wait=refreshDuration
		}
	}}()
	return nil
}
func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		log.Printf("-->unary interceptor: %s",method)
		if interceptor.authMethods[method]{
			return invoker(interceptor.attachToken(ctx),method,req,reply,cc,opts...)
		}
		return invoker(ctx,method,req,reply,cc,opts...)
	}
}
func (interceptor *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		log.Printf("-->stream interceptor: %s",method)
		if interceptor.authMethods[method]{
			return streamer(interceptor.attachToken(ctx),desc,cc,method,opts...)
		}
		return streamer(ctx,desc,cc,method,opts...)
	}
}
func (interceptor *AuthInterceptor) refreshToken() error {
	accessToken, err := interceptor.authClient.Login()
	if err != nil {
		return err
	}
	interceptor.accessToken = accessToken
	log.Printf("token refreshed: %v", accessToken)
	return nil
}

func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	attached:=metadata.AppendToOutgoingContext(ctx,"authorization", interceptor.accessToken)
	return attached
}

func NewAuthInterceptor(authClient *AuthClient, authMethods map[string]bool,refreshDuration time.Duration) (*AuthInterceptor,error) {
	interceptor:= &AuthInterceptor{
		authClient: authClient,
		authMethods: authMethods,
	}

	err:=interceptor.scheduleRefreshToken(refreshDuration)
	if err != nil {
		return nil,err
	}
	return interceptor,nil
}