package notion

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

type QueryDatabaseResult struct {
	Object     string  `json:"object"`
	Results    []Page  `json:"results"`
	NextCursor *string `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
	Type       string  `json:"type"`
	//Page       string  `json:"page"`
}

func (service *Service) QueryDatabase(databaseId string) (*QueryDatabaseResult, *errortools.Error) {
	result := QueryDatabaseResult{}

	requestConfig := go_http.RequestConfig{
		Method:        http.MethodPost,
		Url:           service.url(fmt.Sprintf("databases/%s/query", databaseId)),
		ResponseModel: &result,
	}

	_, _, e := service.httpRequest(&requestConfig)
	if e != nil {
		return nil, e
	}

	return &result, nil
}
