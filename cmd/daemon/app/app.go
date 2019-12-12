package app

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi"
	"log"

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

	cliApi := cliapi.NewCliApi()
	if err := cliApi.Listen(); err != nil {
		return err
	}

	rcv.storage.Start()
	return nil
}

func (rcv *Daemon) Stop() error {
	log.Println("Stopped")
	return nil
}
