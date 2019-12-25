package main

import (
	"github.com/dmytro-kolesnyk/dds/cmd/cli/app"
	"github.com/dmytro-kolesnyk/dds/common/conf"
	"log"
)

func main() {
	configResolver := conf.NewResolver()
	config, err := configResolver.GetConfig()
	if err != nil {
		log.Println("Please specify the config config.yaml")
		return
	}

	cliApp := app.NewApp("localhost", config.CliApi.Port)
	cliApp.Start()
}
