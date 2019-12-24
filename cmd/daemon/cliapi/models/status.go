package models

type StatusResponse struct {
	Status Status `json:"status,omitempty"`
	Error  string `json:"error,omitempty"`
}

type Status struct {
	Uuid     string  `json:"uuid,omitempty"`
	Status   string  `json:"status,omitempty"`
	Progress float32 `json:"progress,omitempty"`
}
