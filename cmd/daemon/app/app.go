package app

import (
	"log"
)

type App interface {
	Start() error
	Stop() error
}

type Daemon struct {
}

func NewDaemon() App {
	return &Daemon{}
}

func (rcv *Daemon) Start() error {
	log.Println("Started")
	return nil
}

func (rcv *Daemon) Stop() error {
	log.Println("Stopped")
	return nil
}
