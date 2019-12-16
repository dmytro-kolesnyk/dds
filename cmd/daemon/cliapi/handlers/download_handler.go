package handlers

import (
	"fmt"
	"net/http"
)

type DownloadHandler struct {
	http.Handler
}

func (rcv *DownloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("download handler")

	response := `{
	  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998"
	}`
	if _, err := fmt.Fprintf(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{}
}
