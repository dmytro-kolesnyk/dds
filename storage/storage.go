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
	lStorage     *localstorage.LocalStorage
	lStoragePath string
	storeLocal   bool
}

// NewStorage function
func NewStorage(config *models.Config) *Storage {
	return &Storage{
		lStorage:     localstorage.NewLocalStorage(),
		splitter:     storage.NewSplitter(config),
		lStoragePath: config.Storage.LocalStoragePath,
		storeLocal:   config.Storage.StoreLocal,
	}
}

// Method used when current node have to distribute file
func (rcv *Storage) SplitSave(data []byte, filename string, strategy string) {
	chanks := rcv.splitter.Split(data, filename, strategy)

	//TODO add logic

}

// Method used when other nodes send their data to distribute
func (rcv *Storage) SaveChunk(chunk storage.Chunk) {
	rcv.lStorage.Save(chunk)
}
