package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/models"
	"github.com/gorilla/mux"
	"net/http"
)

type DeleteHandler struct {
	http.Handler
}

func (rcv *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete handler")

	vars := mux.Vars(r)
	uuid := vars["uuid"]

	// TODO: get real data
	response := &models.DeleteResponse{
		Uuid: uuid,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := fmt.Fprintf(w, string(jsonResponse)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewDeleteHandler() *DeleteHandler {
	return &DeleteHandler{}
}
