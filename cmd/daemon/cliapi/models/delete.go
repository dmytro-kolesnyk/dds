package models

type DeleteResponse struct {
	Uuid  string `json:"uuid,omitempty"`
	Error string `json:"error,omitempty"`
}
