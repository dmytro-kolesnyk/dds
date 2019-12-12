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
}

func NewUploadHandler() *UploadHandler {
	return &UploadHandler{}
}
