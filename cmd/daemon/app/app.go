package app

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi"
	"log"
)

type App interface {
	Start() error
	Stop() error
}

type Daemon struct {
}

func NewDaemon() App {
	return &Daemon{}
}

func (rcv *Daemon) Start() error {
	log.Println("Started")

	cliApi := cliapi.NewCliApi()
	if err := cliApi.Listen(); err != nil {
		return err
	}

	return nil
}

func (rcv *Daemon) Stop() error {
	log.Println("Stopped")
	return nil
}
