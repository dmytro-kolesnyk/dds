package main

import (
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

	app := app.NewDaemon()
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
