package controller

import (
	"fmt"
	"github.com/dmytro-kolesnyk/dds/storage"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Controller struct {
	storage *storage.Storage
}

func NewController(storage *storage.Storage) *Controller {
	return &Controller{storage: storage}
}

func (rcv *Controller) Save(filePath string, strategy string, storeLocally bool) (string, error) {
	if !rcv.isFileExists(filePath) {
		return "", fmt.Errorf("file not found: %v", filePath)
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// TODO: this should return uuid or error
	rcv.storage.Save(bytes, filepath.Base(filePath), strategy)
	return "f4c8de96-4e03-4772-b83c-f8dfbe64e998", nil
}

func (rcv *Controller) isFileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
