package handlers

import (
	"fmt"
	"net/http"
)

type UploadHandler struct {
	http.Handler
}

func (rcv *UploadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("upload handler")

	response := `{
	  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998"
	}`
	if _, err := fmt.Fprintf(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}
