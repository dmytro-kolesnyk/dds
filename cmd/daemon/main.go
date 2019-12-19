package main

import (
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmytro-kolesnyk/dds/cmd/daemon/app"
)

func main() {
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

	fmt.Println("awaiting signal")
	<-sigs
	fmt.Println("exiting")

	if err := app.Stop(); err != nil {
		log.Fatal(err)
	}
}
