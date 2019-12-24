package app

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi"
	"github.com/dmytro-kolesnyk/dds/common/conf"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf/models"
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
	return &Daemon{
		logger:  logger.NewLogger(Daemon{}),
		storage: storage.NewStorage(config),
		cliApi:  cliapi.NewCliApi(config),
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
	rcv.logger.Info("Stopped")

	return nil
}
