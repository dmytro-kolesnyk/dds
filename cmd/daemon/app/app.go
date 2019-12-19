package app

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf"
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
}

func NewDaemon() App {
	return &Daemon{
		storage: storage.NewStorage(),
		logger:  logger.NewLogger(Daemon{}),
	}
}

func (rcv *Daemon) Start() error {
	rcv.logger.Info("Started")

	configResolver := conf.NewResolver()
	config, err := configResolver.GetConfig()
	if err != nil {
		return err
	}

	if err := rcv.storage.Start(); err != nil {
		return err
	}

	cliApi := cliapi.NewCliApi(config)
	if err := cliApi.Listen(); err != nil {
		return err
	}

	return nil
}

func (rcv *Daemon) Stop() error {
	rcv.logger.Info("Stopped")

	return nil
}
