package handlers

import (
	"fmt"
	"net/http"
)

type DeleteHandler struct {
	http.Handler
}

func (rcv *DeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("delete handler")
}

func NewDeleteHandler() *DeleteHandler {
	return &DeleteHandler{}
}
