package serialize

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func JsonToProto(jsonData []byte, message proto.Message) error {
	err := json.Unmarshal(jsonData, message)
	if err != nil {
			return err
	}
	return nil
}

func ProtoToJson(message proto.Message) ([]byte, error) {
	jsonData, err := json.Marshal(message)
	if err != nil {
			return nil, err
	}
	return jsonData, nil
}

func WriteProtoToFile(message proto.Message, filename string) error {
	bytes_, err := proto.Marshal(message)

	if err != nil {
		log.Fatalln("Error marshaling proto message: ", err)
		return err
	}

	if err := os.WriteFile(filename, bytes_, 0644); err != nil {
		log.Fatalln("Error writing to file: ", err)
		return err
	}

	return nil
}

func ReadProtoFromFile(filename string, target proto.Message) error {
	input, err := os.ReadFile(filename)

	if err != nil {
		log.Fatalln("Error reading from file: ", err)
		return err
	}

	if err := proto.Unmarshal(input, target); err != nil {
		log.Fatalln("Error reading from file: ", err)
		return err
	}

	return nil
}

func PackAny(msg proto.Message) *anypb.Any {
	anyMsg, err := anypb.New(msg)
	if err != nil {
			panic(fmt.Errorf("failed to pack message as Any: %v", err))
	}
	return anyMsg
}
