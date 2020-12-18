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

package streams

import (
	"net/http"
	"os"
)

// ServicerGenerated represents the interface for implementing all endpoints for this service
type ServicerGenerated interface {
	/*
		ActivatePipeline - Activates an existing pipeline.
		Parameters:
			id: Pipeline ID
			activatePipelineRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ActivatePipeline(id string, activatePipelineRequest ActivatePipelineRequest, resp ...*http.Response) (*Response, error)
	/*
		Compile - Compiles SPL2 and returns streams JSON.
		Parameters:
			splCompileRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	Compile(splCompileRequest SplCompileRequest, resp ...*http.Response) (*Pipeline, error)
	/*
		CreateConnection - Create a new DSP connection.
		Parameters:
			connectionRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateConnection(connectionRequest ConnectionRequest, resp ...*http.Response) (*ConnectionSaveResponse, error)
	/*
		CreateDataStream - Creates a data stream for a tenant.
		Parameters:
			dataStreamRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateDataStream(dataStreamRequest DataStreamRequest, resp ...*http.Response) (*DataStreamResponse, error)
	/*
		CreatePipeline - Creates a pipeline.
		Parameters:
			pipelineRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreatePipeline(pipelineRequest PipelineRequest, resp ...*http.Response) (*PipelineResponse, error)
	/*
		CreateTemplate - Creates a template for a tenant.
		Parameters:
			templateRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	CreateTemplate(templateRequest TemplateRequest, resp ...*http.Response) (*TemplateResponse, error)
	/*
		DeactivatePipeline - Deactivates an existing pipeline.
		Parameters:
			id: Pipeline ID
			deactivatePipelineRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeactivatePipeline(id string, deactivatePipelineRequest DeactivatePipelineRequest, resp ...*http.Response) (*Response, error)
	/*
		Decompile - Decompiles UPL and returns SPL.
		Parameters:
			decompileRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	Decompile(decompileRequest DecompileRequest, resp ...*http.Response) (*DecompileResponse, error)
	/*
		DeleteConnection - Delete all versions of a connection by its id.
		Parameters:
			connectionId: Connection ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteConnection(connectionId string, resp ...*http.Response) error
	/*
		DeleteDataStream - Deletes a data stream for a tenant.
		Parameters:
			id: ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteDataStream(id string, resp ...*http.Response) error
	/*
		DeleteFile - Delete file.
		Parameters:
			fileId: File ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteFile(fileId string, resp ...*http.Response) error
	/*
		DeleteLookupFile - Delete lookup file.
		Parameters:
			fileId: File ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteLookupFile(fileId string, resp ...*http.Response) error
	/*
		DeletePipeline - Removes a pipeline.
		Parameters:
			id: Pipeline ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeletePipeline(id string, resp ...*http.Response) error
	/*
		DeleteTemplate - Removes a template with a specific ID.
		Parameters:
			templateId: Template ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DeleteTemplate(templateId string, resp ...*http.Response) error
	/*
		DescribeDataStream - Describes a data stream for a tenant.
		Parameters:
			id: ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	DescribeDataStream(id string, resp ...*http.Response) (*DataStreamResponse, error)
	/*
		GetFileMetadata - Get file metadata.
		Parameters:
			fileId: File ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetFileMetadata(fileId string, resp ...*http.Response) (*UploadFileResponse, error)
	/*
		GetFilesMetadata - Returns files metadata.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetFilesMetadata(resp ...*http.Response) (*FilesMetaDataResponse, error)
	/*
		GetInputSchema - Returns the input schema for a function in a pipeline.
		Parameters:
			getInputSchemaRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetInputSchema(getInputSchemaRequest GetInputSchemaRequest, resp ...*http.Response) (*UplType, error)
	/*
		GetLookupFileMetadata - Get lookup file metadata.
		Parameters:
			fileId: File ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetLookupFileMetadata(fileId string, resp ...*http.Response) (*UploadFileResponse, error)
	/*
		GetLookupFilesMetadata - Returns lookup files metadata.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetLookupFilesMetadata(resp ...*http.Response) (*FilesMetaDataResponse, error)
	/*
		GetLookupTable - Returns lookup table results.
		Parameters:
			connectionId: Connection ID
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetLookupTable(connectionId string, query *GetLookupTableQueryParams, resp ...*http.Response) (*LookupTableResponse, error)
	/*
		GetOutputSchema - Returns the output schema for a specified function in a pipeline.
		Parameters:
			getOutputSchemaRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetOutputSchema(getOutputSchemaRequest GetOutputSchemaRequest, resp ...*http.Response) (map[string]UplType, error)
	/*
		GetPipeline - Returns an individual pipeline by version.
		Parameters:
			id: Pipeline ID
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPipeline(id string, query *GetPipelineQueryParams, resp ...*http.Response) (*PipelineResponse, error)
	/*
		GetPipelineLatestMetrics - Returns the latest metrics for a single pipeline.
		Parameters:
			id: Pipeline ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPipelineLatestMetrics(id string, resp ...*http.Response) (*MetricsResponse, error)
	/*
		GetPipelinesStatus - Returns the status of pipelines from the underlying streaming system.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPipelinesStatus(query *GetPipelinesStatusQueryParams, resp ...*http.Response) (*PaginatedResponseOfPipelineJobStatus, error)
	/*
		GetPreviewData - Returns the preview data for a session.
		Parameters:
			previewSessionId: Preview Session ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPreviewData(previewSessionId int64, resp ...*http.Response) (*PreviewData, error)
	/*
		GetPreviewSession - Returns information from a preview session.
		Parameters:
			previewSessionId: Preview Session ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPreviewSession(previewSessionId int64, resp ...*http.Response) (*PreviewState, error)
	/*
		GetPreviewSessionLatestMetrics - Returns the latest metrics for a preview session.
		Parameters:
			previewSessionId: Preview Session ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetPreviewSessionLatestMetrics(previewSessionId int64, resp ...*http.Response) (*MetricsResponse, error)
	/*
		GetRegistry - Returns all functions in JSON format.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetRegistry(query *GetRegistryQueryParams, resp ...*http.Response) (*RegistryModel, error)
	/*
		GetTemplate - Returns an individual template by version.
		Parameters:
			templateId: Template ID
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	GetTemplate(templateId string, query *GetTemplateQueryParams, resp ...*http.Response) (*TemplateResponse, error)
	/*
		ListConnections - Returns a list of connections (latest versions only) by tenant ID.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListConnections(query *ListConnectionsQueryParams, resp ...*http.Response) (*PaginatedResponseOfConnectionResponse, error)
	/*
		ListConnectors - Returns a list of the available connectors.
		Parameters:
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListConnectors(resp ...*http.Response) (*PaginatedResponseOfConnectorResponse, error)
	/*
		ListDataStreams - Returns a list of datastreams for a tenant.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListDataStreams(query *ListDataStreamsQueryParams, resp ...*http.Response) ([]DataStreamResponse, error)
	/*
		ListPipelines - Returns all pipelines.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListPipelines(query *ListPipelinesQueryParams, resp ...*http.Response) (*PaginatedResponseOfPipelineResponse, error)
	/*
		ListTemplates - Returns a list of all templates.
		Parameters:
			query: a struct pointer of valid query parameters for the endpoint, nil to send no query parameters
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ListTemplates(query *ListTemplatesQueryParams, resp ...*http.Response) (*PaginatedResponseOfTemplateResponse, error)
	/*
		PatchPipeline - Patches an existing pipeline.
		Parameters:
			id: Pipeline ID
			pipelinePatchRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PatchPipeline(id string, pipelinePatchRequest PipelinePatchRequest, resp ...*http.Response) (*PipelineResponse, error)
	/*
		PutConnection - Updates an existing DSP connection.
		Parameters:
			connectionId: Connection ID
			connectionPutRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PutConnection(connectionId string, connectionPutRequest ConnectionPutRequest, resp ...*http.Response) (*ConnectionSaveResponse, error)
	/*
		PutTemplate - Updates an existing template.
		Parameters:
			templateId: Template ID
			templatePutRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	PutTemplate(templateId string, templatePutRequest TemplatePutRequest, resp ...*http.Response) (*TemplateResponse, error)
	/*
		ReactivatePipeline - Reactivate a pipeline
		Parameters:
			id: Pipeline ID
			reactivatePipelineRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ReactivatePipeline(id string, reactivatePipelineRequest ReactivatePipelineRequest, resp ...*http.Response) (*PipelineReactivateResponse, error)
	/*
		StartPreview - Creates a preview session for a pipeline.
		Parameters:
			previewSessionStartRequest: Parameters to start a new Preview session
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	StartPreview(previewSessionStartRequest PreviewSessionStartRequest, resp ...*http.Response) (*PreviewStartResponse, error)
	/*
		StopPreview - Stops a preview session.
		Parameters:
			previewSessionId: Preview Session ID
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	StopPreview(previewSessionId int64, resp ...*http.Response) error
	/*
		UpdateConnection - Patches an existing DSP connection.
		Parameters:
			connectionId: Connection ID
			connectionPatchRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateConnection(connectionId string, connectionPatchRequest ConnectionPatchRequest, resp ...*http.Response) (*ConnectionSaveResponse, error)
	/*
		UpdateDataStream - Patches an existing data stream for a tenant.
		Parameters:
			id: ID
			dataStreamRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateDataStream(id string, dataStreamRequest DataStreamRequest, resp ...*http.Response) (*DataStream, error)
	/*
		UpdatePipeline - Updates an existing pipeline.
		Parameters:
			id: Pipeline ID
			pipelineRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdatePipeline(id string, pipelineRequest PipelineRequest, resp ...*http.Response) (*PipelineResponse, error)
	/*
		UpdateTemplate - Patches an existing template.
		Parameters:
			templateId: Template ID
			templatePatchRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UpdateTemplate(templateId string, templatePatchRequest TemplatePatchRequest, resp ...*http.Response) (*TemplateResponse, error)
	/*
		UploadFile - Upload new file.
		Parameters:
			file: Upload file
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	UploadFile(file **os.File, resp ...*http.Response) (*UploadFileResponse, error)
	/*
		ValidatePipeline - Verifies whether the Streams JSON is valid.
		Parameters:
			validateRequest: Request JSON
			resp: an optional pointer to a http.Response to be populated by this method. NOTE: only the first resp pointer will be used if multiple are provided
	*/
	ValidatePipeline(validateRequest ValidateRequest, resp ...*http.Response) (*ValidateResponse, error)
}
