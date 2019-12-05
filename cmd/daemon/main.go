package main

import (
	"dds/cmd/daemon/app"
	"log"
)

func main() {
	app := app.NewDaemon()
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	if err := app.Stop(); err != nil {
		log.Fatal(err)
	}
}
