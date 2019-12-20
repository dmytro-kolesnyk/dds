package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/models"
	"github.com/gorilla/mux"
	"net/http"
)

type DownloadHandler struct {
	http.Handler
}

func (rcv *DownloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("download handler")

	vars := mux.Vars(r)

	request := &models.DownloadRequest{
		Uuid:    vars["uuid"],
		DirPath: vars["dirpath"],
	}

	if err := rcv.Validate(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: get real data
	response := &models.DownloadResponse{
		Uuid: request.Uuid,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := fmt.Fprintf(w, string(jsonResponse)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rcv *DownloadHandler) Validate(request *models.DownloadRequest) error {
	if request.DirPath == "" {
		return fmt.Errorf("no save dir passed")
	}
	return nil
}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{}
}
