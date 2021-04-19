/*
 * Copyright © 2021 Splunk, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"): you may
 * not use this file except in compliance with the License. You may obtain
 * a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
 * License for the specific language governing permissions and limitations
 * under the License.
 *
 * Ingest API
 *
 * Use the Ingest service in Splunk Cloud Services to send event and metrics data, or upload a static file, to Splunk Cloud Services.
 *
 * API version: v1beta2.17 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package ingest

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/splunk/go-dependencies/services"
	"github.com/splunk/go-dependencies/util"
)

const serviceCluster = "api"

type Service services.BaseService

// NewService creates a new ingest service client from the given Config
func NewService(iClient services.IClient) *Service {
	return &Service{Client: iClient}
}

/*
	DeleteAllCollectorTokens - Delete All dsphec tokens for a given tenant.
	Parameters:
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteAllCollectorTokens(resp ...*http.Response) (*map[string]interface{}, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/collector/tokens`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb map[string]interface{}
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	DeleteCollectorToken - Delete dsphec token by name.
	Parameters:
		tokenName
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) DeleteCollectorToken(tokenName string, resp ...*http.Response) (*map[string]interface{}, error) {
	pp := struct {
		TokenName string
	}{
		TokenName: tokenName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/collector/tokens/{{.TokenName}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Delete(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb map[string]interface{}
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	GetCollectorToken - Get the metadata of a dsphec token by name.
	Parameters:
		tokenName
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) GetCollectorToken(tokenName string, resp ...*http.Response) (*HecTokenAccessResponse, error) {
	pp := struct {
		TokenName string
	}{
		TokenName: tokenName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/collector/tokens/{{.TokenName}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb HecTokenAccessResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	ListCollectorTokens - List dsphec tokens for a tenant.
	Parameters:
		query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) ListCollectorTokens(query *ListCollectorTokensQueryParams, resp ...*http.Response) ([]HecTokenAccessResponse, error) {
	values := util.ParseURLParams(query)
	u, err := s.Client.BuildURLFromPathParams(values, serviceCluster, `/ingest/v1beta2/collector/tokens`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Get(services.RequestParams{URL: u})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb []HecTokenAccessResponse
	err = util.ParseResponse(&rb, response)
	return rb, err
}

/*
	PostCollectorTokens - Creates dsphec tokens.
	Parameters:
		hecTokenCreateRequest: The API request schema for the token.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PostCollectorTokens(hecTokenCreateRequest HecTokenCreateRequest, resp ...*http.Response) (*HecTokenCreateResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/collector/tokens`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: hecTokenCreateRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb HecTokenCreateResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	PostEvents - Sends events.
	Parameters:
		event
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PostEvents(event []Event, resp ...*http.Response) (*HttpResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/events`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: event})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb HttpResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	PostMetrics - Sends metric events.
	Parameters:
		metricEvent
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PostMetrics(metricEvent []MetricEvent, resp ...*http.Response) (*HttpResponse, error) {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/metrics`, nil)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Post(services.RequestParams{URL: u, Body: metricEvent})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb HttpResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
	PutCollectorToken - Update the metadata of a dsphec token by name.
	Parameters:
		tokenName
		hecTokenUpdateRequest: The API request schema for the token.
		resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) PutCollectorToken(tokenName string, hecTokenUpdateRequest HecTokenUpdateRequest, resp ...*http.Response) (*HecTokenAccessResponse, error) {
	pp := struct {
		TokenName string
	}{
		TokenName: tokenName,
	}
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/collector/tokens/{{.TokenName}}`, pp)
	if err != nil {
		return nil, err
	}
	response, err := s.Client.Put(services.RequestParams{URL: u, Body: hecTokenUpdateRequest})
	if response != nil {
		defer response.Body.Close()

		// populate input *http.Response if provided
		if len(resp) > 0 && resp[0] != nil {
			*resp[0] = *response
		}
	}
	if err != nil {
		return nil, err
	}
	var rb HecTokenAccessResponse
	err = util.ParseResponse(&rb, response)
	return &rb, err
}

/*
   UploadFiles - Upload a CSV or text file that contains events. The file limit is 1MB or an error is returned.
   Parameters:
       filename
       resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
*/
func (s *Service) UploadFiles(filename string, resp ...*http.Response) error {
	u, err := s.Client.BuildURLFromPathParams(nil, serviceCluster, `/ingest/v1beta2/files`, nil)

	if err != nil {
		return err
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var response *http.Response
	if len(resp) > 0 && resp[0] != nil {
		response = resp[0]
	}

	form := services.FormData{Filename: filepath.Base(filename), Stream: file, Key: "upfile"}

	multipartResp, err := s.Client.Post(services.RequestParams{URL: u, Body: form, Headers: map[string]string{"Content-Type": "multipart/form-data"}})

	if multipartResp != nil {
		defer multipartResp.Body.Close()

		// populate input *http.Response if provided
		if response != nil {
			*response = *multipartResp
		}
	}

	if err != nil {
		return err
	}

	return nil
}
