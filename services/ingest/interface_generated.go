/*
 * Copyright © 2019 Splunk, Inc.
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
 */

// Code generated by gen_interface.go. DO NOT EDIT.

package ingest

import (
	"net/http"
)

// ServicerGenerated represents the interface for implementing all endpoints for this service
type ServicerGenerated interface {
	/*
		DeleteAllCollectorTokens - Delete All dsphec tokens for a given tenant.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteAllCollectorTokens(resp ...*http.Response) (*map[string]interface{}, error)
	/*
		DeleteCollectorToken - Delete dsphec token by name.
		Parameters:
			tokenName
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteCollectorToken(tokenName string, resp ...*http.Response) (*map[string]interface{}, error)
	/*
		GetCollectorToken - Get the metadata of a dsphec token by name.
		Parameters:
			tokenName
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetCollectorToken(tokenName string, resp ...*http.Response) (*HecTokenAccessResponse, error)
	/*
		ListCollectorTokens - List dsphec tokens for a tenant.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListCollectorTokens(query *ListCollectorTokensQueryParams, resp ...*http.Response) ([]HecTokenAccessResponse, error)
	/*
		PostCollectorTokens - Creates dsphec tokens.
		Parameters:
			hecTokenCreateRequest: The API request schema for the token.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PostCollectorTokens(hecTokenCreateRequest HecTokenCreateRequest, resp ...*http.Response) (*HecTokenCreateResponse, error)
	/*
		PostEvents - Sends events.
		Parameters:
			event
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PostEvents(event []Event, resp ...*http.Response) (*HttpResponse, error)
	/*
		PostMetrics - Sends metric events.
		Parameters:
			metricEvent
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PostMetrics(metricEvent []MetricEvent, resp ...*http.Response) (*HttpResponse, error)
	/*
		PutCollectorToken - Update the metadata of a dsphec token by name.
		Parameters:
			tokenName
			hecTokenUpdateRequest: The API request schema for the token.
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PutCollectorToken(tokenName string, hecTokenUpdateRequest HecTokenUpdateRequest, resp ...*http.Response) (*HecTokenAccessResponse, error)
	/*
	   UploadFiles - Upload a CSV or text file that contains events. The file limit is 1MB or an error is returned.
	   Parameters:
	       filename
	       resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UploadFiles(filename string, resp ...*http.Response) error
}
