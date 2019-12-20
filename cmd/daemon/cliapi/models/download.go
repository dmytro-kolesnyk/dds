package models

type DownloadRequest struct {
	Uuid    string
	DirPath string
}

type DownloadResponse struct {
	Uuid  string `json:"uuid,omitempty"`
	Error string `json:"error,omitempty"`
}
