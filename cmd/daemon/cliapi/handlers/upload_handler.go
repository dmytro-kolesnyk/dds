package handlers

import (
	"encoding/json"
	"fmt"
	strategy "github.com/dmytro-kolesnyk/dds/cmd/daemon/app/models"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/models"
	"io/ioutil"
	"net/http"
)

type UploadHandler struct {
	http.Handler
}

func (rcv *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload handler")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request := &models.UploadRequest{}
	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := rcv.Validate(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: get real data
	response := &models.UploadResponse{
		Uuid: "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := fmt.Fprintf(w, string(jsonResponse)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (rcv *UploadHandler) Validate(request *models.UploadRequest) error {
	if request.FilePath == "" {
		return fmt.Errorf("invalid file path: %v", request.FilePath)
	}
	if _, err := strategy.ParseStrategy(request.Strategy); err != nil {
		return err
	}
	return nil
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}
