package fileio

import (
	"fmt"
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
func Write(data []byte, file string) int {
	f, err := os.Open(file)
	check(err)
	// do stuff
	fmt.Println(data)
	f.Close()
	return 0
}
