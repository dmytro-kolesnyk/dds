package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const HttpPrompt string = "http://"
const BaseRoute = "/files"

type Controller struct {
	BaseUrl string
	client  *http.Client
}

func NewController(host, port string) *Controller {
	return &Controller{
		BaseUrl: HttpPrompt + host + ":" + port,
		client:  &http.Client{},
	}
}

func (s *Controller) doRequest(req *http.Request) ([]byte, error) {
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}

//TODO: Add StatusCMDHandler

func (s *Controller) ListCmdHandler(route string) (*ListResp, error) {
	var (
		rawResponse []byte
		response    = &ListResp{}
		err         error
	)
	req, err := http.NewRequest("GET", s.BaseUrl + BaseRoute + route, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for List command, %s\n", err)
	}

	rawResponse, err = s.doRequest(req)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data, %s\n", err)
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse respone, %s\n", err)
	}

	return response, err
}

func (s *Controller) DownloadCmdHandler(route string) (*DownloadResp, error) {
	var (
		rawResponse []byte
		response    = &DownloadResp{}
		err         error
	)
	req, err := http.NewRequest("GET", s.BaseUrl + BaseRoute + route, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for Download command, %s\n", err)
	}

	rawResponse, err = s.doRequest(req)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data, %s\n", err)
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse respone, %s\n", err)
	}

	return response, err
}

func (s *Controller) DeleteCmdHandler(route string) (*DeleteResp, error) {
	var (
		rawResponse []byte
		response    = &DeleteResp{}
		err         error
	)
	req, err := http.NewRequest("DELETE", s.BaseUrl + BaseRoute + route, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request for Delete command, %s\n", err)
	}

	rawResponse, err = s.doRequest(req)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data, %s\n", err)
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse respone, %s\n", err)
	}

	return response, err
}

func (s *Controller) UploadCmdHandler(upload *UploadReq, route string) (*UploadResp, error) {
	var (
		rawResponse []byte
		response    = &UploadResp{}
		err         error
	)
	reqBody, err := json.Marshal(upload)
	if err != nil {
		return nil, fmt.Errorf("failed to create json payload for Upload command, %s\n", err)
	}
	req, err := http.NewRequest("POST", s.BaseUrl + BaseRoute + route, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request for Upload command, %s\n", err)
	}
	rawResponse, err = s.doRequest(req)

	if err != nil {
		return nil, fmt.Errorf("failed to retrieve response data, %s\n", err)
	}

	err = json.Unmarshal(rawResponse, &response)
	if err != nil {
		return nil, fmt.Errorf("failed to parse respone, %s\n", err)
	}

	return response, err
}
