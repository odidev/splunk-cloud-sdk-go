/*
 * Copyright © 2020 Splunk, Inc.
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
 * Collect Service
 *
 * With the Collect service in Splunk Cloud Services, you can manage how data collection jobs ingest event and metric data.
 *
 * API version: v1beta1.8 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package collect

type BaseJob struct {
	// The ID of the connector used in the job.
	ConnectorID string `json:"connectorID"`
	Name        string `json:"name"`
	// The cron schedule, in UTC time format.
	Schedule         string  `json:"schedule"`
	CreateUserID     *string `json:"createUserID,omitempty"`
	CreatedAt        *string `json:"createdAt,omitempty"`
	Id               *string `json:"id,omitempty"`
	LastModifiedAt   *string `json:"lastModifiedAt,omitempty"`
	LastUpdateUserID *string `json:"lastUpdateUserID,omitempty"`
	// Defines whether a job is scheduled or not
	Scheduled *bool   `json:"scheduled,omitempty"`
	Tenant    *string `json:"tenant,omitempty"`
}

type BaseJobPatch struct {
	// The ID of the connector used in the job.
	ConnectorID *string `json:"connectorID,omitempty"`
	// The job name
	Name *string `json:"name,omitempty"`
	// The configuration of the connector used in the job.
	Parameters  map[string]interface{} `json:"parameters,omitempty"`
	ScalePolicy *ScalePolicy           `json:"scalePolicy,omitempty"`
	// The cron schedule, in UTC time format.
	Schedule *string `json:"schedule,omitempty"`
}

// number of jobs deleted.
type DeleteJobsResponse struct {
	Count int32 `json:"count"`
}

type EventExtraField struct {
	// Field name
	Name string `json:"name"`
	// Field value
	Value string `json:"value"`
}

type Execution struct {
	ExecutionUid *string `json:"executionUid,omitempty"`
	JobId        *string `json:"jobId,omitempty"`
	// The given status of the execution
	Status *ExecutionStatus `json:"status,omitempty"`
	Tenant *string          `json:"tenant,omitempty"`
}

// ExecutionStatus : The given status of the execution
type ExecutionStatus string

// List of ExecutionStatus
const (
	ExecutionStatusPending   ExecutionStatus = "pending"
	ExecutionStatusRunning   ExecutionStatus = "running"
	ExecutionStatusCompleted ExecutionStatus = "completed"
	ExecutionStatusAborted   ExecutionStatus = "aborted"
	ExecutionStatusCanceled  ExecutionStatus = "canceled"
)

type ExecutionConflictError struct {
	Code    string    `json:"code"`
	Data    Execution `json:"data"`
	Message string    `json:"message"`
	// The optional details of the error.
	Details map[string]interface{} `json:"details,omitempty"`
	// An optional link to a web page with more information on the error.
	MoreInfo *string `json:"moreInfo,omitempty"`
}

type ExecutionPatch struct {
	// The given status of the execution
	Status ExecutionPatchStatus `json:"status"`
}

// ExecutionPatchStatus : The given status of the execution
type ExecutionPatchStatus string

// List of ExecutionPatchStatus
const (
	ExecutionPatchStatusCanceled ExecutionPatchStatus = "canceled"
)

type Job struct {
	// The ID of the connector used in the job.
	ConnectorID string `json:"connectorID"`
	Name        string `json:"name"`
	// The configuration of the connector used in the job.
	Parameters  map[string]interface{} `json:"parameters"`
	ScalePolicy ScalePolicy            `json:"scalePolicy"`
	// The cron schedule, in UTC time format.
	Schedule         string            `json:"schedule"`
	CreateUserID     *string           `json:"createUserID,omitempty"`
	CreatedAt        *string           `json:"createdAt,omitempty"`
	EventExtraFields []EventExtraField `json:"eventExtraFields,omitempty"`
	Id               *string           `json:"id,omitempty"`
	LastModifiedAt   *string           `json:"lastModifiedAt,omitempty"`
	LastUpdateUserID *string           `json:"lastUpdateUserID,omitempty"`
	// Defines whether a job is scheduled or not
	Scheduled *bool   `json:"scheduled,omitempty"`
	Tenant    *string `json:"tenant,omitempty"`
}

type JobPatch struct {
	// The ID of the connector used in the job.
	ConnectorID      *string           `json:"connectorID,omitempty"`
	EventExtraFields []EventExtraField `json:"eventExtraFields,omitempty"`
	// The job name
	Name *string `json:"name,omitempty"`
	// The configuration of the connector used in the job.
	Parameters  map[string]interface{} `json:"parameters,omitempty"`
	ScalePolicy *ScalePolicy           `json:"scalePolicy,omitempty"`
	// The cron schedule, in UTC time format.
	Schedule *string `json:"schedule,omitempty"`
	// Defines wheather a job is scheduled or not
	Scheduled *bool `json:"scheduled,omitempty"`
}

type JobsPatch struct {
	// The ID of the connector used in the job.
	ConnectorID      *string           `json:"connectorID,omitempty"`
	EventExtraFields []EventExtraField `json:"eventExtraFields,omitempty"`
	ScalePolicy      *ScalePolicy      `json:"scalePolicy,omitempty"`
}

// List of job summaries i.e. scheduling informations, owner, updates, connector.
type ListJobsResponse struct {
	Data []BaseJob `json:"data,omitempty"`
}

// The metadata for the patch jobs operation.
type Metadata struct {
	// The number of jobs that failed to update.
	Failures int64 `json:"failures"`
	// The number of jobs which match the query criteria.
	TotalMatchJobs int64 `json:"totalMatchJobs"`
}

type ModelError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	// The optional details of the error.
	Details map[string]interface{} `json:"details,omitempty"`
	// An optional link to a web page with more information on the error.
	MoreInfo *string `json:"moreInfo,omitempty"`
}

type PatchJobResult struct {
	// The Job ID.
	Id string `json:"id"`
	// Successfully updated or not.
	Updated bool        `json:"updated"`
	Error   *ModelError `json:"error,omitempty"`
}

type PatchJobsResponse struct {
	Data     []PatchJobResult `json:"data"`
	Metadata Metadata         `json:"metadata"`
}

type ScalePolicy struct {
	Static StaticScale `json:"static"`
}

type SingleExecutionResponse struct {
	Data Execution `json:"data"`
}

type SingleJobResponse struct {
	Data Job `json:"data"`
}

type StaticScale struct {
	// The number of collect workers.
	Workers int32 `json:"workers"`
}
