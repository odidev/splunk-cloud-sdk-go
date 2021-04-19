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
	"os"
)

type Error struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}

type Event struct {
	// The raw event content. It can be a string, number, string array, number array, JSON object, map, list, a JSON array, or a byte array.
	Body interface{} `json:"body"`
	// Specifies a JSON object that contains explicit custom fields to be defined at index time.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The host value assigned to the event data. Typically, this is the hostname of the client from which you are sending data.
	Host *string `json:"host,omitempty"`
	// An optional ID that uniquely identifies the event data. It is used to deduplicate the data if same data is set multiple times. If ID is not specified, it will be assigned by the system.
	Id *string `json:"id,omitempty"`
	// Optional nanoseconds part of the timestamp.
	Nanos *int32 `json:"nanos,omitempty"`
	// The source value to assign to the event data. For example, if you are sending data from an app that you are developing, set this key to the name of the app.
	Source *string `json:"source,omitempty"`
	// The sourcetype value assigned to the event data.
	Sourcetype *string `json:"sourcetype,omitempty"`
	// Epoch time in milliseconds.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type FileUploadDetails struct {
	Filename *string `json:"filename,omitempty"`
}

type HecResponse struct {
	// code defines the status of the response. For a complete list of the possible error codes, see https://docs.splunk.com/Documentation/SplunkCloud/latest/Data/TroubleshootHTTPEventCollector#Possible_error_codes
	Code *int64 `json:"code,omitempty"`
	// invalid-event-number defines the index of the first event in error.
	InvalidEventNumber *int64 `json:"invalid-event-number,omitempty"`
	// text describes the response status
	Text *string `json:"text,omitempty"`
}

// GET  /collector/tokens/{tokenName} PUT  /collector/tokens/{tokenName}
type HecTokenAccessResponse struct {
	// ack_enabled is set to true if events sent with the auth token should support indexer acknowledgement  type: bool
	AckEnabled *bool `json:"ack_enabled,omitempty"`
	// allow_query_string_auth is set to true if this token can be passed into the ingest endpoint's query parameter for auth  type: bool
	AllowQueryStringAuth *bool `json:"allow_query_string_auth,omitempty"`
	// created_at is a timestamp that captures when this token was created.  type: string format: date-time
	CreatedAt *string `json:"created_at,omitempty"`
	// created_by is the principal that created the token.  type: string
	CreatedBy *string `json:"created_by,omitempty"`
	// description is an optional description of the token.  type: string
	Description *string `json:"description,omitempty"`
	// disabled is set to true if this auth token has been disabled and cannot be used to send events to HECv1  type: bool
	Disabled *bool `json:"disabled,omitempty"`
	// index is the default value of the index field for records collected using this token.  type: string
	Index *string `json:"index,omitempty"`
	// indexes is a list of index names that this token is allowed to send events to  type: []string
	Indexes []string `json:"indexes,omitempty"`
	// last_modified_at is a timestamp that captures when this token was last modified.  type: string format: date-time
	LastModifiedAt *string `json:"last_modified_at,omitempty"`
	// last_modified_by is the principal that last modified the token.  type: string
	LastModifiedBy *string `json:"last_modified_by,omitempty"`
	// name is the name of the token (unique within the tenant that it belongs to).  type: string
	Name *string `json:"name,omitempty"`
	// source is the default value of the source field for records collected using this token.  type: string
	Source *string `json:"source,omitempty"`
	// sourcetype is the default value of the sourcetype field for records collected using this token.  type: string
	Sourcetype *string `json:"sourcetype,omitempty"`
	// tenant is the tenant that this token belongs to  type: string
	Tenant *string `json:"tenant,omitempty"`
}

// POST  /collector/tokens
type HecTokenCreateRequest struct {
	// name is the name of the token (unique within the tenant that it belongs to).  type: string
	Name string `json:"name"`
	// ack_enabled is set to true if events sent with the auth token should support indexer acknowledgement  type: bool
	AckEnabled *bool `json:"ack_enabled,omitempty"`
	// allow_query_string_auth is set to true if this token can be passed into the ingest endpoint's query parameter for auth  type: bool
	AllowQueryStringAuth *bool `json:"allow_query_string_auth,omitempty"`
	// description is an optional description of the token.  type: string
	Description *string `json:"description,omitempty"`
	// disabled is set to true if this auth token has been disabled and cannot be used to send events to HECv1  type: bool
	Disabled *bool `json:"disabled,omitempty"`
	// index is the default value of the index field for records collected using this token.  type: string
	Index *string `json:"index,omitempty"`
	// indexes is a list of index names that this token is allowed to send events to  type: []string
	Indexes []string `json:"indexes,omitempty"`
	// source is the default value of the source field for records collected using this token.  type: string
	Source *string `json:"source,omitempty"`
	// sourcetype is the default value of the sourcetype field for records collected using this token.  type: string
	Sourcetype *string `json:"sourcetype,omitempty"`
}

// POST /collector/tokens
type HecTokenCreateResponse struct {
	// ack_enabled is set to true if events sent with the auth token should support indexer acknowledgement  type: bool
	AckEnabled *bool `json:"ack_enabled,omitempty"`
	// allow_query_string_auth is set to true if this token can be passed into the ingest endpoint's query parameter for auth  type: bool
	AllowQueryStringAuth *bool `json:"allow_query_string_auth,omitempty"`
	// created_at is a timestamp that captures when this token was created.  type: string format: date-time
	CreatedAt *string `json:"created_at,omitempty"`
	// created_by is the principal that created the token.  type: string
	CreatedBy *string `json:"created_by,omitempty"`
	// description is an optional description of the token.  type: string
	Description *string `json:"description,omitempty"`
	// disabled is set to true if this auth token has been disabled and cannot be used to send events to HECv1  type: bool
	Disabled *bool `json:"disabled,omitempty"`
	// index is the default value of the index field for records collected using this token.  type: string
	Index *string `json:"index,omitempty"`
	// indexes is a list of index names that this token is allowed to send events to  type: []string
	Indexes []string `json:"indexes,omitempty"`
	// last_modified_at is a timestamp that captures when this token was last modified.  type: string format: date-time
	LastModifiedAt *string `json:"last_modified_at,omitempty"`
	// last_modified_by is the principal that last modified the token.  type: string
	LastModifiedBy *string `json:"last_modified_by,omitempty"`
	// name is the name of the token (unique within the tenant that it belongs to).  type: string
	Name *string `json:"name,omitempty"`
	// source is the default value of the source field for records collected using this token.  type: string
	Source *string `json:"source,omitempty"`
	// sourcetype is the default value of the sourcetype field for records collected using this token.  type: string
	Sourcetype *string `json:"sourcetype,omitempty"`
	// tenant is the tenant that this token belongs to.  type: string
	Tenant *string `json:"tenant,omitempty"`
	// token is the token value.  type: string
	Token *string `json:"token,omitempty"`
}

// DELETE /collector/tokens/{tokenName}
type HecTokenDeleteResponse map[string]interface{}

// PUT  /collector/tokens/{tokenName}
type HecTokenUpdateRequest struct {
	// ack_enabled is set to true if events sent with the auth token should support indexer acknowledgement  type: *bool
	AckEnabled *bool `json:"ack_enabled,omitempty"`
	// allow_query_string_auth is set to true if this token can be passed into the ingest endpoint's query parameter for auth  type: *bool
	AllowQueryStringAuth *bool `json:"allow_query_string_auth,omitempty"`
	// description is an optional description of the token.  type: *string
	Description *string `json:"description,omitempty"`
	// disabled is set to true if this auth token has been disabled and cannot be used to send events to HECv1  type: *bool
	Disabled *bool `json:"disabled,omitempty"`
	// index is the default value of the index field for records collected using this token  type: *string
	Index *string `json:"index,omitempty"`
	// indexes is a list of index names that this token is allowed to send events to  type: []string
	Indexes []string `json:"indexes,omitempty"`
	// source is the default value of the source field for records collected using this token  type: *string
	Source *string `json:"source,omitempty"`
	// sourcetype is the default value of the sourcetype field for records collected using this token  type: *string
	Sourcetype *string `json:"sourcetype,omitempty"`
}

type HttpResponse struct {
	Code    *string                `json:"code,omitempty"`
	Details map[string]interface{} `json:"details,omitempty"`
	Message *string                `json:"message,omitempty"`
}

type InlineObject struct {
	Upfile **os.File `json:"upfile,omitempty"`
}

type Metric struct {
	// Name of the metric e.g. CPU, Memory etc.
	Name string `json:"name"`
	// Dimensions allow metrics to be classified e.g. {\"Server\":\"nginx\", \"Region\":\"us-west-1\", ...}
	Dimensions map[string]string `json:"dimensions,omitempty"`
	// Type of metric. Default is g for gauge.
	Type *string `json:"type,omitempty"`
	// Unit of the metric e.g. percent, megabytes, seconds etc.
	Unit *string `json:"unit,omitempty"`
	// Value of the metric. If not specified, it will be defaulted to 0.
	Value *float64 `json:"value,omitempty"`
}

type MetricAttribute struct {
	// Optional. If set, individual metrics inherit these dimensions and can override any and/or all of them.
	DefaultDimensions map[string]string `json:"defaultDimensions,omitempty"`
	// Optional. If set, individual metrics inherit this type and can optionally override.
	DefaultType *string `json:"defaultType,omitempty"`
	// Optional. If set, individual metrics inherit this unit and can optionally override.
	DefaultUnit *string `json:"defaultUnit,omitempty"`
}

type MetricEvent struct {
	// Specifies multiple related metrics e.g. Memory, CPU etc.
	Body       []Metric         `json:"body"`
	Attributes *MetricAttribute `json:"attributes,omitempty"`
	// The host value assigned to the event data. Typically, this is the hostname of the client from which you are sending data.
	Host *string `json:"host,omitempty"`
	// An optional ID that uniquely identifies the metric data. It is used to deduplicate the data if same data is set multiple times. If ID is not specified, it will be assigned by the system.
	Id *string `json:"id,omitempty"`
	// Optional nanoseconds part of the timestamp.
	Nanos *int32 `json:"nanos,omitempty"`
	// The source value to assign to the event data. For example, if you are sending data from an app that you are developing, set this key to the name of the app.
	Source *string `json:"source,omitempty"`
	// The sourcetype value assigned to the event data.
	Sourcetype *string `json:"sourcetype,omitempty"`
	// Epoch time in milliseconds.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

type UploadSuccessResponse struct {
	Code    *string            `json:"code,omitempty"`
	Details *FileUploadDetails `json:"details,omitempty"`
	Message *string            `json:"message,omitempty"`
}
