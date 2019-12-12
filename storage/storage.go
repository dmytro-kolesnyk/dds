package storage

import (
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
func (rcv *Storage) Start() {
	rcv.lStorage.Save()
}
