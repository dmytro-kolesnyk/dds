package models

type ListResponse struct {
	Files []File `json:"files,omitempty"`
	Error string `json:"error,omitempty"`
}

type File struct {
	Uuid     string `json:"uuid,omitempty"`
	FileName string `json:"file_name,omitempty"`
}
