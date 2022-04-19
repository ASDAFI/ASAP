package grpc_server

import (
	"farm/configs"
	"farm/src/infrastructure/controllers"
	"farm/src/middlewares"
	"farm/src/proto/service/farm"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func newFarmServer() controllers.FarmServer {
	return controllers.FarmServer{}
}

func RunInsecureGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", configs.Config.Server.Host,
		configs.Config.Server.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var kaep = keepalive.EnforcementPolicy{
		MinTime: 5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		// PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}

	var kasp = keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
		// MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
		MaxConnectionAgeGrace: 3 * time.Second, // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
		Time:                  1 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
		Timeout:               1 * time.Second, // Wait 1 second for the ping ack before assuming the connection is dead
	}
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_auth.StreamServerInterceptor(middlewares.Authenticate),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(middlewares.Authenticate),
		)),
		grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp),
	)

	farmServer := newFarmServer()
	farm.RegisterFarmServer(grpcServer, &farmServer)
	handleGRPCServerSignals(grpcServer)
	log.Fatal(grpcServer.Serve(lis))
}

func handleGRPCServerSignals(grpcServer *grpc.Server) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT)
	signal.Notify(c, syscall.SIGKILL)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		for range c {
			grpcServer.GracefulStop()
			os.Exit(0)
		}
	}()
}
