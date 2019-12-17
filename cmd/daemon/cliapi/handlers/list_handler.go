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

	response := `{
	   "files": [{
		  "uuid": "f4c8de96-4e03-4772-b83c-f8dfbe64e998",
		  "file_name": "file1.avi"
		}]
	}`
	if _, err := fmt.Fprintf(w, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func NewListHandler() *ListHandler {
	return &ListHandler{}
}
