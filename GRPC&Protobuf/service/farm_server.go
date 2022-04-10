package service

import (
	ASAP "ASAP/pb/proto"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type ConnectionServer struct {
	ASAP.UnimplementedConnectionServiceServer
}

func NewConnectionServer() *ConnectionServer {
	return &ConnectionServer{
		ASAP.UnimplementedConnectionServiceServer{},

	}
}

func (server *ConnectionServer) Connect(context context.Context, request *ASAP.CreateConnectionRequest) (*ASAP.CreateConnectionResponse, error){
	connection := request.GetMessageRequest()
	log.Printf("receive connection request : %v ", connection.Request)
	if connection.GetRequest()==""{
		log.Print("could not make connection")
		return nil, fmt.Errorf("could not make connection")
	}
	if err := contextError(context); err != nil {
		return nil, err
	}

	response := &ASAP.CreateConnectionResponse{MessageResponse: connection}
	return response, nil

}
func contextError(ctx context.Context) error{
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}

}
func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}