package storage

import (
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	communicationServer "github.com/dmytro-kolesnyk/dds/communication_server"
	"github.com/dmytro-kolesnyk/dds/storage/localstorage"
	storage "github.com/dmytro-kolesnyk/dds/storage/splitter"
	"log"
)

// Storage struct
type Storage struct {
	splitter     *storage.Splitter
	offset       int
	lStorage     *localstorage.LocalStorage
	cServer      *communicationServer.CommunicationServer
	lStoragePath string
	storeLocal   bool
}

// NewStorage function
func NewStorage(config *models.Config) *Storage {
	return &Storage{
		lStorage:   localstorage.NewLocalStorage(config),
		splitter:   storage.NewSplitter(config),
		cServer:    communicationServer.NewCommunicationServer(config),
		storeLocal: config.Storage.StoreLocal,
		offset:     config.Storage.Offset,
	}
}

// Start method
func (rcv *Storage) Start() error {
	if err := rcv.cServer.Start(); err != nil {
		log.Println("searching for neighbors")
		return err
	}

	return nil
}

// Method used when current node have to distribute file
func (rcv *Storage) Save(data []byte, filename string, strategy string, offset int) {
	log.Printf("Save file %s with strategy %s and offset %d", filename, strategy, offset)
	chunks := rcv.splitter.Split(data, filename, strategy, offset)

	for _, c := range chunks {
		rcv.saveChunk(c)
	}
	//TODO add logic to handle table
}

// Method used when other nodes send their data to distribute
func (rcv *Storage) saveChunk(chunk storage.Chunk) {
	rcv.lStorage.Save(chunk)
}
