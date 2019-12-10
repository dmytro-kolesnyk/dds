package cliapi

import (
	"./handlers"
	"github.com/gorilla/mux"
	"net/http"
)

type CliApi struct {
}

func (rcv *CliApi) start() error {
	router := mux.NewRouter()

	router.Handle("/files", handlers.NewListHandler()).Methods(http.MethodGet)
	router.Handle("/files", handlers.NewUploadHandler()).Methods(http.MethodPost)
	router.Handle("/files/{uuid}", handlers.NewDownloadHandler()).Methods(http.MethodGet)
	router.Handle("/files/{uuid}", handlers.NewDeleteHandler()).Methods(http.MethodDelete)
	router.Handle("/files/{uuid}/status", handlers.NewStatusHandler()).Methods(http.MethodGet)

	return http.ListenAndServe(":8080", router)
}

func NewCliApi() *CliApi {
	return &CliApi{}
}
