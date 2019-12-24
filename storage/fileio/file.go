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

// Read function
func Read(file string, offset int, lenght int) []byte {
	f, err := os.Open(file)
	check(err)
	// do stuff
	f.Close()
	b := make([]byte, 2)
	b[0] = 2
	b[1] = 2
	return b
}

// Write function
func Write(data []byte, path string) int {
	// Open a new file for writing only
	file, err := os.OpenFile(
		path,
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0666,
	)
	check(err)
	defer file.Close()
	// do stuff

	bytesWritten, err := file.Write(data)
	check(err)

	//
	log.Printf("Wrote %d bytes. into path %s\n", bytesWritten, path)
	return bytesWritten
}
