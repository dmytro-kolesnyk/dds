package storage

import (
	"github.com/dmytro-kolesnyk/dds/localstorage"
)

// Storage struct
type Storage struct {
	localStorage *localstorage.LocalStorage
}

// NewStorage function
func NewStorage() *Storage {
	return &Storage{
		localStorage: localstorage.NewLocalStorage(),
	}
}

// Start method
func (rcv *Storage) Start() {
	rcv.localStorage.Save()
}
