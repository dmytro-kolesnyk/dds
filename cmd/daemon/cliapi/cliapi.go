package cliapi

import (
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type CliApi struct {
}

func (rcv *CliApi) Listen() error {
	router := mux.NewRouter()

	router.Handle("/files", handlers.NewListHandler()).Methods(http.MethodGet)
	router.Handle("/files", handlers.NewUploadHandler()).Methods(http.MethodPost)
	router.Handle("/files/{uuid}", handlers.NewDownloadHandler()).Methods(http.MethodGet)
	router.Handle("/files/{uuid}", handlers.NewDeleteHandler()).Methods(http.MethodDelete)
	router.Handle("/files/{uuid}/status", handlers.NewStatusHandler()).Methods(http.MethodGet)

	// TODO: port must be configured
	return http.ListenAndServe(":8080", router)
}

func NewCliApi() *CliApi {
	return &CliApi{}
}
