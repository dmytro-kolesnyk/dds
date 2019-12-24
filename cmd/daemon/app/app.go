package app

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf/models"
	"log"

	communicationServer "github.com/dmytro-kolesnyk/dds/communication_server"
	"github.com/dmytro-kolesnyk/dds/discovery"
	"github.com/dmytro-kolesnyk/dds/node"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf"
	"github.com/dmytro-kolesnyk/dds/common/logger"
	"github.com/dmytro-kolesnyk/dds/storage"
	"github.com/google/uuid"
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

	if err := rcv.cliApi.Listen(); err != nil {
		return err
	}

	if err := rcv.storage.Start(); err != nil {
		return err
	}

	return nil
}

func (rcv *Daemon) Stop() error {
	rcv.logger.Info("Stopped")

	return nil
}
