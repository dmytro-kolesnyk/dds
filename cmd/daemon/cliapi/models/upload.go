package models

type UploadRequest struct {
	FilePath     string `json:"file_path"`
	Strategy     string `json:"strategy"`
	StoreLocally bool   `json:"store_locally"`
}

type UploadResponse struct {
	Uuid  string `json:"uuid,omitempty"`
	Error string `json:"error,omitempty"`
}
