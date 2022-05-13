package notion

import (
	"fmt"
	"net/http"
	"net/url"

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

func (service *Service) QueryDatabase(databaseId string) (*[]Page, *errortools.Error) {
	pages := []Page{}

	values := url.Values{}

	for {
		result := QueryDatabaseResult{}

		requestConfig := go_http.RequestConfig{
			Method:        http.MethodPost,
			Url:           service.url(fmt.Sprintf("databases/%s/query?%s", databaseId, values.Encode())),
			ResponseModel: &result,
		}

		_, _, e := service.httpRequest(&requestConfig)
		if e != nil {
			return nil, e
		}

		pages = append(pages, result.Results...)

		if !result.HasMore || result.NextCursor == nil {
			break
		}

		values.Set("start_cursor", *result.NextCursor)
	}

	return &pages, nil
}
