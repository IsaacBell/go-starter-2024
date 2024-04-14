package main

import _ "github.com/joho/godotenv/autoload"

import (
	"log"
	"flag"

	"github.com/Soapstone-Services/go-template-2024/pkg/api"

	"github.com/Soapstone-Services/go-template-2024/pkg/utl/config"
	errorUtils "github.com/Soapstone-Services/go-template-2024/pkg/utl/errors"

	basic "github.com/Soapstone-Services/go-template-2024/protogen/basic"
)

func checkProtobufHeartbeat() {
	helloWorld := basic.Hello {
		Name: "Hello World!",
	}

	log.Println("===============================================")
	log.Println(&helloWorld)
	log.Println("===============================================")
}

func main() {
	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	errorUtils.CheckErr(err)

	checkProtobufHeartbeat()
	
	errorUtils.CheckErr(api.Start(cfg))
}

