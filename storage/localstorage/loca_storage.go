package localstorage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf/models"
	"github.com/dmytro-kolesnyk/dds/storage/fileio"
	"github.com/dmytro-kolesnyk/dds/storage/splitter"
	"log"
)

// LocalStorage contains actual data
type LocalStorage struct {
	lStoragePath string
	// Some Table to keep track of storedData
}

// NewLocalStorage method
func NewLocalStorage(config *models.Config) *LocalStorage {
	return &LocalStorage{
		lStoragePath: config.Storage.LocalStoragePath,
	}
}

// Save method
func (rcv LocalStorage) Save(chunk storage.Chunk) {
	log.Printf("Save chunk of the file %s with id %d and date %s, maxId = %d", chunk.FileName, chunk.Id, chunk.Date, chunk.MaxId)
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(chunk)
	if err != nil {
		fmt.Println(`failed gob Encode`, err)
	}

	fileio.Write(b.Bytes(), rcv.lStoragePath)

	// someUsefulInfo := fileio.Write(chunk.Data, rcv.path)
	// Populate table with someUsefulInfo + uuid-id-filename
}
