package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message)  (string){
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers: false,
	Indent: " ",
	UseProtoNames: true,

	}
	return marshaler.Format(message)
}