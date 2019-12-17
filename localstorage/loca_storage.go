package localstorage

import (
	"github.com/dmytro-kolesnyk/dds/fileio"
)

// LocalStorage contains actual data
type LocalStorage struct {
	data []byte
	path string
}

// Save method
func (rcv LocalStorage) Save() {
	fileio.Write(rcv.data, rcv.path)
}

// NewLocalStorage method
func NewLocalStorage() *LocalStorage {
	return &LocalStorage{}
}
