package app

import (
	"log"
	"os"
	"strconv"

	communicationServer "github.com/dmytro-kolesnyk/dds/communication_server"
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

	//cliApi := cliapi.NewCliApi()
	//if err := cliApi.Listen(); err != nil {
	//	return err
	//}

	//defer discoverer.Stop()

	cs := communicationServer.NewCommunicationServer()
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return err
	}

	if err := cs.Start(port); err != nil {
		log.Println("searching for neighbors")
		// log.Fatalln(err)
		return err
	}

	return nil
}

func (rcv *Daemon) Stop() error {
	log.Println("Stopped")
	return nil
}
