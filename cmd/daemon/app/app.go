package app

import (
	"log"

	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf"
	"github.com/dmytro-kolesnyk/dds/storage"
)

type App interface {
	Start() error
	Stop() error
}

type Daemon struct {
	storage *storage.Storage
}

func NewDaemon() App {
	return &Daemon{
		storage: storage.NewStorage(),
	}
}

func (rcv *Daemon) Start() error {
	log.Println("Started")

	configResolver := conf.NewResolver()
	config, err := configResolver.GetConfig()
	if err != nil {
		return err
	}

	cliApi := cliapi.NewCliApi(config)
	cliApi.Listen()

	if err := rcv.storage.Start(); err != nil {
		return err
	}

	return nil
}

func (rcv *Daemon) Stop() error {
	log.Println("Stopped")
	return nil
}
