package localstorage

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	"github.com/dmytro-kolesnyk/dds/storage/fileio"
	"github.com/dmytro-kolesnyk/dds/storage/splitter"
	"log"
)

// LocalStorage contains actual data
type LocalStorage struct {
	lStoragePath string
	TrackTable   *trackTable
}

// NewLocalStorage method
func NewLocalStorage(config *models.Config) *LocalStorage {
	filePath := config.Storage.LocalStoragePath
	return &LocalStorage{
		lStoragePath: filePath,
		TrackTable:   NewTrackTable(fileio.Size(filePath)),
	}
}

// For now, this table will be not stored locally.
// Need to think how to store consistently and not trigger disk for each write
// -1 create timed flushing to disk
//    then we need to cut file and probably lose data
type trackTable struct {
	chunkPointers []chunkPointer
	pointer       int64 // = file.size
}

// TODO add init method for reload app (we need populate chunkPointers after refresh)
func NewTrackTable(pointer int64) *trackTable {
	return &trackTable{
		pointer:       pointer,
		chunkPointers: make([]chunkPointer, 0, 0),
	}
}

type chunkPointer struct {
	uuid   string
	index  int
	offset int64
	size   int
}

func newChunkPointer(chunk storage.Chunk, offset int64, size int) *chunkPointer {
	return &chunkPointer{
		uuid:   chunk.Uuid,
		index:  chunk.Id,
		offset: offset,
		size:   size,
	}
}

func (rcv *trackTable) addChunk(chunk storage.Chunk, writtenBytes int) {
	currentPointer := rcv.pointer
	chunkPointer := newChunkPointer(chunk, currentPointer, writtenBytes)
	rcv.chunkPointers = append(rcv.chunkPointers, *chunkPointer)

	log.Printf("Add to tracked table pointer with uuid = %s, with start position = %d and size = %d", chunkPointer.uuid, chunkPointer.offset, writtenBytes)

	// Or we can just use file.size()
	rcv.pointer += int64(writtenBytes)
}

// Save method
func (rcv LocalStorage) Save(chunk storage.Chunk) {
	log.Printf("Save chunk of the fileName = %s with id = %d and uuid = %s, maxId = %d", chunk.FileName, chunk.Id, chunk.Uuid, chunk.MaxId)
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	err := encoder.Encode(chunk)
	if err != nil {
		fmt.Println(`Failed gob Encode`, err)
		// TODO Panic probably
		return
	}

	writtenBytes := fileio.Write(b.Bytes(), rcv.lStoragePath)
	rcv.TrackTable.addChunk(chunk, writtenBytes)
}

// This method will be refactored. Probably all logic, need to be in separate track_table.go file
func (rcv LocalStorage) GetChunk(uuid string, index int) *storage.Chunk {
	path := rcv.lStoragePath
	for _, c := range rcv.TrackTable.chunkPointers {
		if c.uuid == uuid && c.index == index {
			cBytes := fileio.Read(path, c.offset, c.size)
			return decodeChunk(cBytes)
		}
	}
	return nil
}

func decodeChunk(cBytes []byte) *storage.Chunk {
	var chunk storage.Chunk
	buf := bytes.NewBuffer(cBytes)
	dec := gob.NewDecoder(buf)
	_ = dec.Decode(&chunk)
	return &chunk
}
