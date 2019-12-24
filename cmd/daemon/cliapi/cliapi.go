package cliapi

import (
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/handlers"
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type CliApi struct {
	port int
}

func (rcv *CliApi) Listen() {
	address := fmt.Sprintf(":%d", rcv.port)

	router := mux.NewRouter()
	router.Methods(http.MethodGet).Path("/files").Handler(handlers.NewListHandler())
	router.Methods(http.MethodPost).Path("/files").Handler(handlers.NewUploadHandler())
	router.Methods(http.MethodGet).Path("/files/{uuid}").Queries("dirpath", "{dirpath}").Handler(handlers.NewDownloadHandler())
	router.Methods(http.MethodDelete).Path("/files/{uuid}").Handler(handlers.NewDeleteHandler())
	router.Methods(http.MethodGet).Path("/files/{uuid}/status").Handler(handlers.NewStatusHandler())

	go func() {
		log.Fatal(http.ListenAndServe(address, router))
	}()
}

func NewCliApi(config *models.Config) *CliApi {
	return &CliApi{port: config.CliApi.Port}
}
