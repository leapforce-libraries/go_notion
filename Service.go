package notion

import (
	"fmt"
	"net/http"

	errortools "github.com/leapforce-libraries/go_errortools"
	go_http "github.com/leapforce-libraries/go_http"
)

const (
	apiName        string = "Notion"
	apiUrl         string = "https://api.notion.com/v1"
	defaultVersion string = "2022-02-22"
)

type ServiceConfig struct {
	BearerToken string
	Version     *string
}

type Service struct {
	bearerToken   string
	version       string
	httpService   *go_http.Service
	errorResponse *ErrorResponse
}

// methods
//
func NewService(serviceConfig *ServiceConfig) (*Service, *errortools.Error) {
	if serviceConfig == nil {
		return nil, errortools.ErrorMessage("ServiceConfig must not be a nil pointer")
	}

	service := Service{
		bearerToken: serviceConfig.BearerToken,
		version:     defaultVersion,
	}

	if serviceConfig.Version != nil {
		service.version = *serviceConfig.Version
	}

	httpServiceConfig := go_http.ServiceConfig{}
	httpService, e := go_http.NewService(&httpServiceConfig)
	if e != nil {
		return nil, e
	}
	service.httpService = httpService
	return &service, nil
}

func (service *Service) httpRequest(requestConfig *go_http.RequestConfig) (*http.Request, *http.Response, *errortools.Error) {
	header := (*requestConfig).NonDefaultHeaders
	if header == nil {
		header = &http.Header{}
	}

	// add authentication header
	header.Set("Authorization", fmt.Sprintf("Bearer %s", service.bearerToken))
	// add version header
	header.Set("Notion-Version", service.version)
	// add version header
	header.Set("Content-Type", "application/json")

	(*requestConfig).NonDefaultHeaders = header

	service.errorResponse = &ErrorResponse{}
	requestConfig.ErrorModel = service.errorResponse

	request, response, e := service.httpService.HttpRequest(requestConfig)
	if e != nil {
		if service.errorResponse.Message != "" {
			e.SetMessage(service.errorResponse.Message)
		}
	}

	return request, response, e
}

func (service *Service) url(path string) string {
	return fmt.Sprintf("%s/%s", apiUrl, path)
}

func (service *Service) ApiName() string {
	return apiName
}

func (service *Service) ApiKey() string {
	return service.bearerToken[:10]
}

func (service *Service) ApiCallCount() int64 {
	return service.httpService.RequestCount()
}

func (service *Service) ApiReset() {
	service.httpService.ResetRequestCount()
}

func (service *Service) ErrorResponse() *ErrorResponse {
	return service.errorResponse
}
