package handlers

import (
	"fmt"
	"net/http"
)

type ListHandler struct {
	http.Handler
}

func (rcv *ListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("list handler")
}

func NewListHandler() *ListHandler {
	return &ListHandler{}
}
