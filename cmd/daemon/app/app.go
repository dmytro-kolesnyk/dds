package app

import (
	"log"
	"os"
	"strconv"

	communicationServer "github.com/dmytro-kolesnyk/dds/communication_server"
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

	/*
		configResolver := conf.NewResolver()
		config, err := configResolver.GetConfig()
		if err != nil {
			return err
		}

		cliApi := cliapi.NewCliApi(config)
		if err := cliApi.Listen(); err != nil {
			return err
		}
	*/

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}
	cs := communicationServer.NewCommunicationServer(port)

	if err := cs.Start(); err != nil {
		log.Println("searching for neighbors")
		return err
	}

	return nil
}

func (rcv *Daemon) Stop() error {
	log.Println("Stopped")
	return nil
}
