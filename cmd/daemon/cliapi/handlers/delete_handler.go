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

	response := `{
	  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998"
	}`
	if _, err := fmt.Fprintf(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewDeleteHandler() *DeleteHandler {
	return &DeleteHandler{}
}
