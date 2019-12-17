package cliapi

import (
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/handlers"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/conf/models"
	"github.com/gorilla/mux"
	"net/http"
)

type CliApi struct {
	port int
}

func (rcv *CliApi) Listen() error {
	router := mux.NewRouter()

	router.Handle("/files", handlers.NewListHandler()).Methods(http.MethodGet)
	router.Handle("/files", handlers.NewUploadHandler()).Methods(http.MethodPost)
	router.Handle("/files/{uuid}", handlers.NewDownloadHandler()).Methods(http.MethodGet)
	router.Handle("/files/{uuid}", handlers.NewDeleteHandler()).Methods(http.MethodDelete)
	router.Handle("/files/{uuid}/status", handlers.NewStatusHandler()).Methods(http.MethodGet)

	return http.ListenAndServe(fmt.Sprintf(":%d", rcv.port), router)
}

func NewCliApi(config *models.Config) *CliApi {
	return &CliApi{port: config.CliApi.Port}
}
