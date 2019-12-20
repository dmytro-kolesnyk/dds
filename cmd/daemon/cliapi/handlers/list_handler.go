package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dmytro-kolesnyk/dds/cmd/daemon/cliapi/models"
	"net/http"
)

type ListHandler struct {
	http.Handler
}

func (rcv *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list handler")

	files := make([]models.File, 0)

	// TODO: get real data
	files = append(files, models.File{
		Uuid:     "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
		FileName: "file1.avi",
	}, models.File{
		Uuid:     "cafc5637-e1e7-4cda-ae38-b56862c4c387",
		FileName: "file2.avi",
	})
	response := models.ListResponse{Files: files}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if _, err := fmt.Fprintf(w, string(jsonResponse)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewListHandler() *ListHandler {
	return &ListHandler{}
}
