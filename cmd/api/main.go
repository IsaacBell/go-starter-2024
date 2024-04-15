package main

import (
	"flag"
	"log"

	"github.com/Soapstone-Services/go-template-2024/pkg/api"
	"github.com/Soapstone-Services/go-template-2024/pkg/api/serialize"
	"github.com/Soapstone-Services/go-template-2024/pkg/utl/config"
	_ "github.com/joho/godotenv/autoload"

	errorUtils "github.com/Soapstone-Services/go-template-2024/pkg/utl/errors"

	basic "github.com/Soapstone-Services/go-template-2024/protogen/basic"
	"google.golang.org/protobuf/types/known/anypb"
)

// todo - move this to the "serialize" package
func checkProtobufHeartbeat() {
	helloWorld := basic.Hello {
		Name: "Hello World!",
	}

	log.Println("===============================================")
	log.Println("protobuffers are working: ", &helloWorld)

	emailChannel := &basic.EmailChannel{
		Email: "user@example.com",
	}

	var emailChannelAsAny *anypb.Any
	emailChannelAsAny, err := anypb.New(emailChannel)
	if err != nil {
		log.Fatalf("Failed to create Any message: %v", err)
	}

	u := basic.User {
		Id: 0,
		Username: "Jane Doe",
		IsActive: true,
		Email: "user@example.com",
		CommunicationChannels: []*anypb.Any{
			serialize.PackAny(emailChannelAsAny),
		},
	}

	log.Println("---")
	log.Println("example user data: ", &u)
	log.Println("---")
	json, _ := serialize.ProtoToJson(&u)
	log.Println("example serialized json: ", json)
	log.Println("===============================================")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	errorUtils.CheckErr(err)

	checkProtobufHeartbeat()

	errorUtils.CheckErr(api.Start(cfg))
}

