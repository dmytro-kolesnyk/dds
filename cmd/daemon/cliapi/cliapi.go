package cliapi

import (
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/handlers"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/controller"
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CliApi struct {
	config     *models.Config
	controller *controller.Controller
}

func (rcv *CliApi) Listen() {
	address := fmt.Sprintf(":%d", rcv.config.CliApi.Port)

	router := mux.NewRouter()
	router.Methods(http.MethodGet).Path("/files").Handler(handlers.NewListHandler())
	router.Methods(http.MethodPost).Path("/files").Handler(handlers.NewUploadHandler(rcv.config, rcv.controller))
	router.Methods(http.MethodGet).Path("/files/{uuid}").Queries("dirpath", "{dirpath}").Handler(handlers.NewDownloadHandler())
	router.Methods(http.MethodDelete).Path("/files/{uuid}").Handler(handlers.NewDeleteHandler())
	router.Methods(http.MethodGet).Path("/files/{uuid}/status").Handler(handlers.NewStatusHandler())

	go func() {
		log.Fatal(http.ListenAndServe(address, router))
	}()
}

func NewCliApi(config *models.Config, controller *controller.Controller) *CliApi {
	return &CliApi{config: config, controller: controller}
}
