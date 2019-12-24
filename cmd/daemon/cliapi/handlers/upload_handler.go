package handlers

import (
	"encoding/json"
	"fmt"
	strategy "github.com/dmytro-kolesnyk/dds/cmd/daemon/app/models"
	cliApiModels "github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/models"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/controller"
	"github.com/dmytro-kolesnyk/dds/common/conf/models"
	"io/ioutil"
	"net/http"
)

type UploadHandler struct {
	http.Handler
	config     *models.Config
	controller *controller.Controller
}

func (rcv *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload handler")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	request := &cliApiModels.UploadRequest{}
	if err := json.Unmarshal(body, &request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := rcv.Validate(request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uuid, err := rcv.controller.Save(request.FilePath, request.Strategy, request.StoreLocally)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) // TODO: distinct user vs server errors
		return
	}
	response := &cliApiModels.UploadResponse{
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

func (rcv *UploadHandler) Validate(request *cliApiModels.UploadRequest) error {
	if request.FilePath == "" {
		return fmt.Errorf("invalid file path: %v", request.FilePath)
	}
	if _, err := strategy.ParseStrategy(request.Strategy); err != nil {
		return err
	}
	return nil
}

func NewUploadHandler(config *models.Config, controller *controller.Controller) *UploadHandler {
	return &UploadHandler{config: config, controller: controller}
}
