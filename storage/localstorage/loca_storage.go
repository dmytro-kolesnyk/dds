package localstorage

import (
	"github.com/dmytro-kolesnyk/dds/storage/fileio"
	"github.com/dmytro-kolesnyk/dds/storage/splitter"
)

// LocalStorage contains actual data
type LocalStorage struct {
	// Some Table to keep track of storedData
}

// NewLocalStorage method
func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}

// SplitSave method
func (rcv LocalStorage) Save(chunk storage.Chunk) {
	// Probably we can use information from Chunk to store it more efficiently (uuid-id-filename table with offsets to search)
	fileio.Write(chunk.Data, rcv.path)
	//someUsefulInfo := fileio.Write(chunk.Data, rcv.path)
	// Populate table with someUsefulInfo + uuid-id-filename
}
