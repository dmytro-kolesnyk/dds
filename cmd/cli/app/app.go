package app

import (
	"github.com/dmytro-kolesnyk/dds/cmd/cli/command_line"
	"github.com/dmytro-kolesnyk/dds/cmd/cli/controller"
)

type App struct {
	host string
	port int
}

func NewApp(host string, port int) *App {
	return &App{host, port}
}

func (rcv *App) Start() {

	restClient := controller.NewController(rcv.host, rcv.port)

	command_line.Init(restClient)

}

func (rcv *App) Stop() error {
	return nil
}
