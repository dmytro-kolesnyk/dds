package fileio

import (
	"log"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Size(file string) int64 {

	open, err := os.Open(file)
	if os.IsNotExist(err) {
		return 0
	} else {
		check(err)
	}

	info, err := open.Stat()
	check(err)

	return info.Size()
}

// Read function
func Read(filePath string, offset int64, size int) []byte {
	file, err := os.OpenFile(
		filePath,
		os.O_RDONLY,
		0666,
	)
	defer file.Close()
	check(err)

	buffer := make([]byte, size)

	rBytes, err := file.ReadAt(buffer, offset)
	check(err)

	if rBytes != size {
		println("%d != size %d ", rBytes, size)
	}

	return buffer
}

// Write function
func Write(data []byte, path string) int {
	// Open a new file for writing only
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0666,
	)
	defer file.Close()
	check(err)

	bytesWritten, err := file.Write(data)
	check(err)

	log.Printf("Wrote %d bytes. into path %s\n", bytesWritten, path)
	return bytesWritten
}
