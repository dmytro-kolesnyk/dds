package main

import (
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmytro-kolesnyk/dds/cmd/daemon/app"
	"github.com/dmytro-kolesnyk/dds/common/logger"
)

func main() {
	logger := logger.NewMainLogger()
	logger.Info("starting daemon...")

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	configResolver := conf.NewResolver()
	config, err := configResolver.GetConfig()
	if err != nil {
		log.Println("Please specify the config config.yaml")
		return
	}

	app := app.NewDaemon(config)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	logger.Info("awaiting signal")
	<-sigs
	logger.Info("exiting")

	if err := app.Stop(); err != nil {
		log.Fatal(err)
	}
}
