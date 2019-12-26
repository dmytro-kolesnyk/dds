package app

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/controller"
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	"github.com/dmytro-kolesnyk/dds/common/logger"
	"github.com/dmytro-kolesnyk/dds/storage"
)

type App interface {
	Start() error
	Stop() error
}

type Daemon struct {
	storage *storage.Storage
	logger  *logger.Logger
	cliApi  *cliapi.CliApi
}

func NewDaemon(config *models.Config) App {
	store := storage.NewStorage(config)
	return &Daemon{
		logger:  logger.NewLogger(Daemon{}),
		storage: store,
		cliApi:  cliapi.NewCliApi(config, controller.NewController(store)),
	}
}

func (rcv *Daemon) Start() error {
	rcv.logger.Info("Started")

	rcv.cliApi.Listen()
	if err := rcv.storage.Start(); err != nil {
		return err
	}

	return nil
}

func (rcv *Daemon) Stop() error {
	if err := rcv.storage.Stop(); err != nil {
		return err
	}
	rcv.logger.Info("Stopped")

	return nil
}
