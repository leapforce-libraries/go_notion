package notion

import (
	"fmt"
	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
	"net/http"
)

type QueryDatabaseResult struct {
	Object     string  `json:"object"`
	Results    []Page  `json:"results"`
	NextCursor *string `json:"next_cursor"`
	HasMore    bool    `json:"has_more"`
	Type       string  `json:"type"`
	//Page       string  `json:"page"`
}

func (service *Service) QueryDatabase(databaseId string) (*[]Page, *errortools.Error) {
	pages := []Page{}

	body := struct {
		StartCursor string `json:"start_cursor"`
	}{}

	for {
		result := QueryDatabaseResult{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodPost,
			Url:           service.url(fmt.Sprintf("databases/%s/query", databaseId)),
			ResponseModel: &result,
		}

		if body.StartCursor != "" {
			requestConfig.BodyModel = body
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		pages = append(pages, result.Results...)

		if !result.HasMore || result.NextCursor == nil {
			break
		}

		body.StartCursor = *result.NextCursor
	}

	return &pages, nil
}
