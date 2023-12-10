package main

import _ "github.com/joho/godotenv/autoload"

import (
	"flag"

	"github.com/Soapstone-Services/go-template-2024/pkg/api"

	"github.com/Soapstone-Services/go-template-2024/pkg/utl/config"
	errorUtils "github.com/Soapstone-Services/go-template-2024/pkg/utl/errors"
)

func main() {
	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	errorUtils.CheckErr(err)

	errorUtils.CheckErr(api.Start(cfg))
}

