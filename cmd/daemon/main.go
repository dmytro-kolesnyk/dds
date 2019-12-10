package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dmytro-kolesnyk/dds/cmd/daemon/app"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	app := app.NewDaemon()
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

	if err := app.Stop(); err != nil {
		log.Fatal(err)
	}
}
