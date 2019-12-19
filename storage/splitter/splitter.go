package storage

import (
	"time"
)

// Splitter struct
type Splitter struct {
	defaultStrategy string
}

func NewSplitter() *Splitter {
	return &Splitter{}
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
		Uuid:     uuid,
		Id:       id,
		MaxId:    maxId,
		FileName: fileName,
		Data:     data,
		Date:     date,
	}
}

//We assume that all chunks have same size
func (rcv *Splitter) Split(data []byte, fileName string, strategy string, offset int) []Chunk {
	// TODO use strategy and offset
	time := time.Now()
	maxId := 3

	chunks := make([]Chunk, maxId)
	// Simple impl
	for i := 0; i < maxId; i++ {
		uuid := "rand-" + string(i)
		chunks[i] = *NewChunk(uuid, i, maxId, fileName, data, time)
	}

	return chunks
}
