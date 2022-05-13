package notion

import (
	"encoding/json"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	n_types "github.com/leapforce-libraries/go_notion/types"
)

type Page struct {
	Object         string                 `json:"object"`
	Id             string                 `json:"id"`
	CreatedTime    n_types.DateTimeString `json:"created_time"`
	LastEditedTime n_types.DateTimeString `json:"last_edited_time"`
	CreatedBy      struct {
		Object string `json:"object"`
		Id     string `json:"id"`
	} `json:"created_by"`
	LastEditedBy struct {
		Object string `json:"object"`
		Id     string `json:"id"`
	} `json:"last_edited_by"`
	Cover  json.RawMessage `json:"cover"`
	Icon   json.RawMessage `json:"icon"`
	Parent struct {
		Type       string `json:"type"`
		DatabaseId string `json:"database_id"`
	} `json:"parent"`
	Archived   bool            `json:"archived"`
	Properties json.RawMessage `json:"properties"`
}

func (service *Service) CreatePage(page *Page) (*Page, *errortools.Error) {
	resultPage := Page{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url("pages"),
		BodyModel:     page,
		ResponseModel: &resultPage,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &resultPage, nil
}

func (service *Service) UpdatePage(page *Page) (*Page, *errortools.Error) {
	resultPage := Page{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPatch,
		Url:           service.url("pages"),
		BodyModel:     page,
		ResponseModel: &resultPage,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &resultPage, nil
}
