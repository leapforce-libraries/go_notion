package notion

import (
	"encoding/json"
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	n_types "github.com/leapforce-libraries/go_notion/types"
)

type PageParent struct {
	Type       string `json:"type"`
	DatabaseId string `json:"database_id"`
}

type Page struct {
	Object         *string                 `json:"object,omitempty"`
	Id             *string                 `json:"id,omitempty"`
	CreatedTime    *n_types.DateTimeString `json:"created_time,omitempty"`
	LastEditedTime *n_types.DateTimeString `json:"last_edited_time,omitempty"`
	CreatedBy      *struct {
		Object string `json:"object"`
		Id     string `json:"id"`
	} `json:"created_by,omitempty"`
	LastEditedBy *struct {
		Object string `json:"object"`
		Id     string `json:"id"`
	} `json:"last_edited_by,omitempty"`
	Cover      json.RawMessage `json:"cover,omitempty"`
	Icon       json.RawMessage `json:"icon,omitempty"`
	Parent     *PageParent     `json:"parent,omitempty"`
	Archived   *bool           `json:"archived,omitempty"`
	Properties json.RawMessage `json:"properties,omitempty"`
}

type TextContent struct {
	Content string  `json:"content"`
	Link    *string `json:"link"`
}

type Text struct {
	Type        *string     `json:"type,omitempty"`
	Text        TextContent `json:"text"`
	Annotations *struct {
		Bold          bool   `json:"bold"`
		Italic        bool   `json:"italic"`
		Strikethrough bool   `json:"strikethrough"`
		Underline     bool   `json:"underline"`
		Code          bool   `json:"code"`
		Color         string `json:"color"`
	} `json:"annotations,omitempty"`
	PlainText *string `json:"plain_text,omitempty"`
	Href      *string `json:"href,omitempty"`
}

type FieldText struct {
	Id       string `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	RichText []Text `json:"rich_text"`
}

type FieldTitle struct {
	Id    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Title []Text `json:"title"`
}

type FieldPhoneNumber struct {
	Id          string  `json:"id,omitempty"`
	Type        string  `json:"type,omitempty"`
	PhoneNumber *string `json:"phone_number"`
}

type FieldEmail struct {
	Id    string  `json:"id,omitempty"`
	Type  string  `json:"type,omitempty"`
	Email *string `json:"email"`
}

type FieldUrl struct {
	Id   string  `json:"id,omitempty"`
	Type string  `json:"type,omitempty"`
	Url  *string `json:"url"`
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

func (service *Service) UpdatePage(pageId string, page *Page) (*Page, *errortools.Error) {
	resultPage := Page{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPatch,
		Url:           service.url(fmt.Sprintf("pages/%s", pageId)),
		BodyModel:     page,
		ResponseModel: &resultPage,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &resultPage, nil
}
