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

	response := `{
	  "status": [
		{
		  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
		  "status": "downloading",
		  "progress": 31.5
		}
	  ]
	}`
	if _, err := fmt.Fprintf(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewStatusHandler() *StatusHandler {
	return &StatusHandler{}
}
