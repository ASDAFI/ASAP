package client

import (
	ASAP "ASAP/pb/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

type ConnectionClient struct {
	service ASAP.ConnectionServiceClient
}

func NewConnectionClient(cc *grpc.ClientConn) *ConnectionClient {
	service:= ASAP.NewConnectionServiceClient(cc)
	return &ConnectionClient{service: service}
}
func (client *ConnectionClient)Connect(message *ASAP.Connection){
	request:= &ASAP.CreateConnectionRequest{MessageRequest: message}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response,err:=client.service.Connect(context.WithValue(ctx,"authorization","ali"),request)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			log.Print("connection already exists")
		} else {
			log.Fatal("cannot create connection", err)

		}
		return
	}
	log.Print(response.MessageResponse)

}
