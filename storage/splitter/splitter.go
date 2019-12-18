package storage

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf/models"
	"time"
)

// Splitter struct
type Splitter struct {
	defaultStrategy string
}

func NewSplitter(config *models.Config) *Splitter {
	return &Splitter{
		defaultStrategy: config.Storage.LocalStoragePath,
	}
}

type Chunk struct {
	Uuid     string
	Id       int
	MaxId    int
	FileName string
	Data     []byte
	Date     time.Time
}

func NewChunk(uuid string, id int, maxId int, fileName string, data []byte, date Time) *Chunk {
	return &Chunk{
		Uuid:     uuid,
		Id:       id,
		MaxId:    maxId,
		FileName: fileName,
		Data:     data,
		Date:     date,
	}
}

func (rcv *Splitter) Split(data []byte, fileName string, strategy string) []Chunk {
	// TODO use strategy
	time := time.Now()
	maxId := 3

	chunks := make([]Chunk, maxId)
	// Simple impl
	for i := 1; i < maxId; i++ {
		uuid := "rand-" + string(i)
		NewChunk(uuid, i, maxId, fileName, data, time)
	}

	return chunks
}
