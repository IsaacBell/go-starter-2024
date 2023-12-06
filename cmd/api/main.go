package main

import (
	"flag"

	"github.com/Soapstone-Services/go-template-2024/pkg/api"

	"github.com/Soapstone-Services/go-template-2024/pkg/utl/config"
)

func main() {
	cfgPath := flag.String("p", "./conf.local.yaml", "Path to config file")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	checkErr(err)

	checkErr(api.Start(cfg))
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
