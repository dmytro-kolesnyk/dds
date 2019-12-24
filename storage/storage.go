package storage

import (
	"log"
	"os"
	"strconv"

	communicationServer "github.com/dmytro-kolesnyk/dds/communication_server"
	"github.com/dmytro-kolesnyk/dds/localstorage"
)

// Storage struct
type Storage struct {
	lStorage *localstorage.LocalStorage
}

// NewStorage function
func NewStorage() *Storage {
	return &Storage{
		lStorage: localstorage.NewLocalStorage(),
	}
}

// Start method
func (rcv *Storage) Start() error {
	port, err := strconv.Atoi(os.Getenv("PORT")) // [FIXME] read from conf.yaml
	if err != nil {
		return err
	}

	// [TODO] move this to "Storage" fields
	cs := communicationServer.NewCommunicationServer(port)

	if err := cs.Start(); err != nil {
		log.Println("searching for neighbors")
		return err
	}

	return nil
}
