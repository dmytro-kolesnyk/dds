package storage

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf/models"
	"github.com/dmytro-kolesnyk/dds/storage/localstorage"
	storage "github.com/dmytro-kolesnyk/dds/storage/splitter"
)

// Storage struct
type Storage struct {
	splitter *storage.Splitter
	// Communication_Server
	// AllNodes []  // Probably need to be updated (health check)
	// All nodes (external and itself) represented by this struct.
	// It mb difficult to create same logic when we including current node, which will be used when storeLocal flag is true
	offset       int
	lStorage     *localstorage.LocalStorage
	lStoragePath string
	storeLocal   bool
}

// NewStorage function
func NewStorage(config *models.Config) *Storage {
	return &Storage{
		lStorage:   localstorage.NewLocalStorage(config),
		splitter:   storage.NewSplitter(),
		storeLocal: config.Storage.StoreLocal,
		offset:     config.Storage.Offset,
	}
}

// Method used when current node have to distribute file
func (rcv *Storage) Save(data []byte, filename string, strategy string, offset int) {
	chanks := rcv.splitter.Split(data, filename, strategy, offset)

	for _, c := range chanks {
		rcv.saveChunk(c)
	}
	//TODO add logic
}

// Method used when other nodes send their data to distribute
func (rcv *Storage) saveChunk(chunk storage.Chunk) {
	rcv.lStorage.Save(chunk)
}
