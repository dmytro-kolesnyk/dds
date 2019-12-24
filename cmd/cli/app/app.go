package app

import (
	"github.com/dmytro-kolesnyk/dds/cmd/cli/command_line"
	"github.com/dmytro-kolesnyk/dds/cmd/cli/controller"
)

type App struct {
	host string
	port string
}

func NewApp(host, port string) *App {
	return &App{host, port}
}

func (a *App) Start(){

	restClient := controller.NewController(a.host, a.port)

	command_line.Init(restClient)

}

func (a *App) Stop() error {
	return nil
}

