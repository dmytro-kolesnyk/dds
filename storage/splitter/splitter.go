package storage

import (
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	"log"
	"time"
)

// Splitter struct
type Splitter struct {
	defaultStrategy string
}

func NewSplitter(config *models.Config) *Splitter {
	return &Splitter{
		defaultStrategy: config.Storage.DefaultStrategy,
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

func NewChunk(uuid string, id int, maxId int, fileName string, data []byte, date time.Time) *Chunk {
	return &Chunk{
		Uuid:     uuid, // Alias for fileName (collision evasion)
		Id:       id,
		MaxId:    maxId,
		FileName: fileName,
		Data:     data,
		Date:     date,
	}
}

//We assume that all chunks have same size
func (rcv *Splitter) Split(data []byte, fileName string, strategy string) []Chunk {
	log.Printf("Split file = %s with strategy = %s", fileName, strategy)
	// TODO use strategy and offset
	createdTime := time.Now()
	maxId := 3
	uuid := fileName + "-uuid"

	chunks := make([]Chunk, maxId)
	// Simple impl
	for i := 0; i < maxId; i++ {
		chunks[i] = *NewChunk(uuid, i, maxId, fileName, data, createdTime)
	}

	return chunks
}
