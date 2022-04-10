package serializer

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
)

func WriteProtobufToJSON(message proto.Message, filename string) error  {
	data := ProtobufToJSON(message)
	//the function used in the video was deprecated and the new ones format is a bit different than in the video
	err := ioutil.WriteFile(filename,[]byte(data),0644)
	if err != nil {
		return fmt.Errorf("Cannot json data to file: %w", err)
	}
	return nil
}

func WriteProtobufToBinaryFile(message proto.Message,filename string)  error{
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto messafe to binary: %w", err)
		}
		err = ioutil.WriteFile(filename,data,0644)
	if err != nil {
		return fmt.Errorf("Cannot binary data to file: %w", err)
	}
	return nil
}
func ReadProtobufFromFile(message proto.Message,filename string)  error{
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Cannot read binary data from file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("Cannot unmarshal binary to proto message: %w", err)
	}
	return nil
}