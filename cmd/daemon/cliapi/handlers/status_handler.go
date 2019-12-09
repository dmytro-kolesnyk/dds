package handlers

import (
	"fmt"
	"net/http"
)

type StatusHandler struct {
	http.Handler
}

func (rcv *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("status handler")
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}
