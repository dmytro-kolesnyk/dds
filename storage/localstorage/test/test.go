package main

import (
	"bytes"
	"encoding/gob"
	"github.com/dmytro-kolesnyk/dds/common/conf"
	"github.com/dmytro-kolesnyk/dds/storage"
)

const fileName = "TEST_FILE"
const uuid = "TEST_FILE-uuid"
const TEST_DATA = "ы11яфы232фывasd31asd2351521r12rfszc1ff1g"

type File struct {
	Uuid string
	Data []byte
}

func NewFile(uuid string, data []byte) *File {
	return &File{
		Uuid: uuid,
		Data: data,
	}
}

func main() {
	chunk := NewFile(TEST_DATA, make([]byte, 5))
	testSave(chunk)
}

func testSave(chunk *File) {
	b := bytes.Buffer{}
	encoder := gob.NewEncoder(&b)
	encoder.Encode(chunk)

	config, _ := conf.NewResolver().GetConfig()
	newStorage := storage.NewStorage(config)
	newStorage.Save(b.Bytes(), fileName, "test-strategy")

	testRead(newStorage)
}

func testRead(storage *storage.Storage) {

	chunk := storage.Read(uuid, 1)

	println(chunk.Id == 1)
	println(chunk.Uuid == uuid)
	println(chunk.FileName == fileName)

	var file File
	buf := bytes.NewBuffer(chunk.Data)
	dec := gob.NewDecoder(buf)
	_ = dec.Decode(&file)

	println(file.Uuid == TEST_DATA)
}
