package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/models"
	"github.com/gorilla/mux"
	"net/http"
)

type StatusHandler struct {
	http.Handler
}

func (rcv *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("status handler")

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// TODO: get real data
	response := &models.StatusResponse{
		Status: models.Status{
			Uuid:     uuid,
			Status:   "complete",
			Progress: 100.0,
		},
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := fmt.Fprintf(w, string(jsonResponse)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}
