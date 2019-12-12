package app

import (
	"log"

	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi"
	communicationServer "github.com/dmytro-kolesnyk/dds/communication_server"
	"github.com/dmytro-kolesnyk/dds/discovery"
	"github.com/dmytro-kolesnyk/dds/node"
	"github.com/google/uuid"
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

	discoverer := discovery.NewDiscovery(
		uuid.New().String(),
		"_dds._tcp",
		"local.",
		3451,
	)

	neighbours := make(chan *node.Node)
	if err := discoverer.Start(neighbours); err != nil {
		//log.Fatalln(err)
		return err
	}
	//defer discoverer.Stop()

	cs := communicationServer.NewCommunicationServer()
	if err := cs.Start(":3451", neighbours); err != nil {
		// log.Fatalln(err)
		return err
	}

	return nil
}

func (rcv *Daemon) Stop() error {
	log.Println("Stopped")
	return nil
}
