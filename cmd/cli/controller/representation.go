package controller

type UploadReq struct {
	FilePath     string `json:"file_path"`
	Strategy     string `json:"strategy"`
	StoreLocally bool   `json:"store_locally"`
}

type UploadResp struct {
	UUID string `json:"uuid"`
}

type DownloadResp struct {
	UUID string `json:"uuid"`
}

type DeleteResp struct {
	UUID string `json:"uuid"`
}

type ListResp struct {
	Files []*File `json:"files"`
}

type File struct {
	UUID     string `json:"uuid"`
	FileName string `json:"file_name"`
}

//TODO: Add Status command representation