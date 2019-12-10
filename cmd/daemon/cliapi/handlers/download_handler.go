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
}

func NewDownloadHandler() *DownloadHandler {
	return &DownloadHandler{}
}
