// Package streams -- generated by scloudgen
// !! DO NOT EDIT !!
//
package streams

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/pkg/streams"
)

// activatePipeline -- Activates an existing pipeline.
var activatePipelineCmd = &cobra.Command{
	Use:   "activate-pipeline",
	Short: "Activates an existing pipeline.",
	RunE:  impl.ActivatePipeline,
}

// compile -- Compiles SPL2 and returns streams JSON.
var compileCmd = &cobra.Command{
	Use:   "compile",
	Short: "Compiles SPL2 and returns streams JSON.",
	RunE:  impl.Compile,
}

// createConnection -- Create a new DSP connection.
var createConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Create a new DSP connection.",
	RunE:  impl.CreateConnection,
}

// createPipeline -- Creates a pipeline.
var createPipelineCmd = &cobra.Command{
	Use:   "create-pipeline",
	Short: "Creates a pipeline.",
	RunE:  impl.CreatePipeline,
}

// createTemplate -- Creates a template for a tenant.
var createTemplateCmd = &cobra.Command{
	Use:   "create-template",
	Short: "Creates a template for a tenant.",
	RunE:  impl.CreateTemplate,
}

// deactivatePipeline -- Deactivates an existing pipeline.
var deactivatePipelineCmd = &cobra.Command{
	Use:   "deactivate-pipeline",
	Short: "Deactivates an existing pipeline.",
	RunE:  impl.DeactivatePipeline,
}

// decompile -- Decompiles UPL and returns SPL.
var decompileCmd = &cobra.Command{
	Use:   "decompile",
	Short: "Decompiles UPL and returns SPL.",
	RunE:  impl.Decompile,
}

// deleteConnection -- Delete all versions of a connection by its id.
var deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Delete all versions of a connection by its id.",
	RunE:  impl.DeleteConnection,
}

// deleteFile -- Delete file.
var deleteFileCmd = &cobra.Command{
	Use:   "delete-file",
	Short: "Delete file.",
	RunE:  impl.DeleteFile,
}

// deleteLookupFile -- Delete lookup file.
var deleteLookupFileCmd = &cobra.Command{
	Use:   "delete-lookup-file",
	Short: "Delete lookup file.",
	RunE:  impl.DeleteLookupFile,
}

// deletePipeline -- Removes a pipeline.
var deletePipelineCmd = &cobra.Command{
	Use:   "delete-pipeline",
	Short: "Removes a pipeline.",
	RunE:  impl.DeletePipeline,
}

// deleteTemplate -- Removes a template with a specific ID.
var deleteTemplateCmd = &cobra.Command{
	Use:   "delete-template",
	Short: "Removes a template with a specific ID.",
	RunE:  impl.DeleteTemplate,
}

// getFileMetadata -- Get file metadata.
var getFileMetadataCmd = &cobra.Command{
	Use:   "get-file-metadata",
	Short: "Get file metadata.",
	RunE:  impl.GetFileMetadata,
}

// getFilesMetadata -- Returns files metadata.
var getFilesMetadataCmd = &cobra.Command{
	Use:   "get-files-metadata",
	Short: "Returns files metadata.",
	RunE:  impl.GetFilesMetadata,
}

// getInputSchema -- Returns the input schema for a function in a pipeline.
var getInputSchemaCmd = &cobra.Command{
	Use:   "get-input-schema",
	Short: "Returns the input schema for a function in a pipeline.",
	RunE:  impl.GetInputSchema,
}

// getLookupFileMetadata -- Get lookup file metadata.
var getLookupFileMetadataCmd = &cobra.Command{
	Use:   "get-lookup-file-metadata",
	Short: "Get lookup file metadata.",
	RunE:  impl.GetLookupFileMetadata,
}

// getLookupFilesMetadata -- Returns lookup files metadata.
var getLookupFilesMetadataCmd = &cobra.Command{
	Use:   "get-lookup-files-metadata",
	Short: "Returns lookup files metadata.",
	RunE:  impl.GetLookupFilesMetadata,
}

// getLookupTable -- Returns lookup table results.
var getLookupTableCmd = &cobra.Command{
	Use:   "get-lookup-table",
	Short: "Returns lookup table results.",
	RunE:  impl.GetLookupTable,
}

// getOutputSchema -- Returns the output schema for a specified function in a pipeline.
var getOutputSchemaCmd = &cobra.Command{
	Use:   "get-output-schema",
	Short: "Returns the output schema for a specified function in a pipeline.",
	RunE:  impl.GetOutputSchema,
}

// getPipeline -- Returns an individual pipeline by version.
var getPipelineCmd = &cobra.Command{
	Use:   "get-pipeline",
	Short: "Returns an individual pipeline by version.",
	RunE:  impl.GetPipeline,
}

// getPipelineLatestMetrics -- Returns the latest metrics for a single pipeline.
var getPipelineLatestMetricsCmd = &cobra.Command{
	Use:   "get-pipeline-latest-metrics",
	Short: "Returns the latest metrics for a single pipeline.",
	RunE:  impl.GetPipelineLatestMetrics,
}

// getPipelinesStatus -- Returns the status of pipelines from the underlying streaming system.
var getPipelinesStatusCmd = &cobra.Command{
	Use:   "get-pipelines-status",
	Short: "Returns the status of pipelines from the underlying streaming system.",
	RunE:  impl.GetPipelinesStatus,
}

// getPreviewData -- Returns the preview data for a session.
var getPreviewDataCmd = &cobra.Command{
	Use:   "get-preview-data",
	Short: "Returns the preview data for a session.",
	RunE:  impl.GetPreviewData,
}

// getPreviewSession -- Returns information from a preview session.
var getPreviewSessionCmd = &cobra.Command{
	Use:   "get-preview-session",
	Short: "Returns information from a preview session.",
	RunE:  impl.GetPreviewSession,
}

// getPreviewSessionLatestMetrics -- Returns the latest metrics for a preview session.
var getPreviewSessionLatestMetricsCmd = &cobra.Command{
	Use:   "get-preview-session-latest-metrics",
	Short: "Returns the latest metrics for a preview session.",
	RunE:  impl.GetPreviewSessionLatestMetrics,
}

// getRegistry -- Returns all functions in JSON format.
var getRegistryCmd = &cobra.Command{
	Use:   "get-registry",
	Short: "Returns all functions in JSON format.",
	RunE:  impl.GetRegistry,
}

// getTemplate -- Returns an individual template by version.
var getTemplateCmd = &cobra.Command{
	Use:   "get-template",
	Short: "Returns an individual template by version.",
	RunE:  impl.GetTemplate,
}

// listConnections -- Returns a list of connections (latest versions only) by tenant ID.
var listConnectionsCmd = &cobra.Command{
	Use:   "list-connections",
	Short: "Returns a list of connections (latest versions only) by tenant ID.",
	RunE:  impl.ListConnections,
}

// listConnectors -- Returns a list of the available connectors.
var listConnectorsCmd = &cobra.Command{
	Use:   "list-connectors",
	Short: "Returns a list of the available connectors.",
	RunE:  impl.ListConnectors,
}

// listPipelines -- Returns all pipelines.
var listPipelinesCmd = &cobra.Command{
	Use:   "list-pipelines",
	Short: "Returns all pipelines.",
	RunE:  impl.ListPipelines,
}

// listTemplates -- Returns a list of all templates.
var listTemplatesCmd = &cobra.Command{
	Use:   "list-templates",
	Short: "Returns a list of all templates.",
	RunE:  impl.ListTemplates,
}

// patchPipeline -- Patches an existing pipeline.
var patchPipelineCmd = &cobra.Command{
	Use:   "patch-pipeline",
	Short: "Patches an existing pipeline.",
	RunE:  impl.PatchPipeline,
}

// putConnection -- Updates an existing DSP connection.
var putConnectionCmd = &cobra.Command{
	Use:   "put-connection",
	Short: "Updates an existing DSP connection.",
	RunE:  impl.PutConnection,
}

// putTemplate -- Updates an existing template.
var putTemplateCmd = &cobra.Command{
	Use:   "put-template",
	Short: "[not implemented] Updates an existing template.",
	RunE:  impl.PutTemplate,
}

// reactivatePipeline -- Reactivate a pipeline
var reactivatePipelineCmd = &cobra.Command{
	Use:   "reactivate-pipeline",
	Short: "Reactivate a pipeline",
	RunE:  impl.ReactivatePipeline,
}

// startPreview -- Creates a preview session for a pipeline.
var startPreviewCmd = &cobra.Command{
	Use:   "start-preview",
	Short: "Creates a preview session for a pipeline.",
	RunE:  impl.StartPreview,
}

// stopPreview -- Stops a preview session.
var stopPreviewCmd = &cobra.Command{
	Use:   "stop-preview",
	Short: "Stops a preview session.",
	RunE:  impl.StopPreview,
}

// updateConnection -- Patches an existing DSP connection.
var updateConnectionCmd = &cobra.Command{
	Use:   "update-connection",
	Short: "Patches an existing DSP connection.",
	RunE:  impl.UpdateConnection,
}

// updatePipeline -- Updates an existing pipeline.
var updatePipelineCmd = &cobra.Command{
	Use:   "update-pipeline",
	Short: "Updates an existing pipeline.",
	RunE:  impl.UpdatePipeline,
}

// updateTemplate -- Patches an existing template.
var updateTemplateCmd = &cobra.Command{
	Use:   "update-template",
	Short: "Patches an existing template.",
	RunE:  impl.UpdateTemplate,
}

// uploadFile -- Upload new file.
var uploadFileCmd = &cobra.Command{
	Use:   "upload-file",
	Short: "Upload new file.",
	RunE:  impl.UploadFile,
}

// uploadLookupFile -- Upload new lookup file.
var uploadLookupFileCmd = &cobra.Command{
	Use:   "upload-lookup-file",
	Short: "Upload new lookup file.",
	RunE:  impl.UploadLookupFile,
}

// validatePipeline -- Verifies whether the Streams JSON is valid.
var validatePipelineCmd = &cobra.Command{
	Use:   "validate-pipeline",
	Short: "Verifies whether the Streams JSON is valid.",
	RunE:  impl.ValidatePipeline,
}

func init() {
	streamsCmd.AddCommand(activatePipelineCmd)

	var activatePipelineId string
	activatePipelineCmd.Flags().StringVar(&activatePipelineId, "id", "", "This is a required parameter. Pipeline ID")
	activatePipelineCmd.MarkFlagRequired("id")

	var activatePipelineActivateLatestVersion string
	activatePipelineCmd.Flags().StringVar(&activatePipelineActivateLatestVersion, "activate-latest-version", "false", "Set to true to activate the latest version of the pipeline. Set to false to use the previously activated version of the pipeline. Defaults to true.")

	var activatePipelineAllowNonRestoredState string
	activatePipelineCmd.Flags().StringVar(&activatePipelineAllowNonRestoredState, "allow-non-restored-state", "false", "Set to true to allow the pipeline to ignore any unused progress states. In some cases, when a data pipeline is changed, the progress state will be stored for functions that no longer exist, so this must be set to activate a pipeline in this state. Defaults to false.")

	var activatePipelineSkipRestoreState string
	activatePipelineCmd.Flags().StringVar(&activatePipelineSkipRestoreState, "skip-restore-state", "false", "Set to true to start reading from the latest input rather than from where the pipeline's previous run left off, which can cause data loss. Defaults to false.")

	streamsCmd.AddCommand(compileCmd)

	var compileInputDatafile string
	compileCmd.Flags().StringVar(&compileInputDatafile, "input-datafile", "", "The SPL2 representation of a pipeline or function parameters")

	var compileValidate string
	compileCmd.Flags().StringVar(&compileValidate, "validate", "false", "A boolean flag to indicate whether the pipeline should be validated.")

	streamsCmd.AddCommand(createConnectionCmd)

	var createConnectionConnectorId string
	createConnectionCmd.Flags().StringVar(&createConnectionConnectorId, "connector-id", "", "This is a required parameter. The ID of the parent connector.")
	createConnectionCmd.MarkFlagRequired("connector-id")

	var createConnectionData string
	createConnectionCmd.Flags().StringVar(&createConnectionData, "data", "", "This is a required parameter. The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.")
	createConnectionCmd.MarkFlagRequired("data")

	var createConnectionDescription string
	createConnectionCmd.Flags().StringVar(&createConnectionDescription, "description", "", "This is a required parameter. The description of the connection.")
	createConnectionCmd.MarkFlagRequired("description")

	var createConnectionName string
	createConnectionCmd.Flags().StringVar(&createConnectionName, "name", "", "This is a required parameter. The name of the connection.")
	createConnectionCmd.MarkFlagRequired("name")

	streamsCmd.AddCommand(createPipelineCmd)

	var createPipelineName string
	createPipelineCmd.Flags().StringVar(&createPipelineName, "name", "", "This is a required parameter. The name of the pipeline.")
	createPipelineCmd.MarkFlagRequired("name")

	var createPipelineBypassValidation string
	createPipelineCmd.Flags().StringVar(&createPipelineBypassValidation, "bypass-validation", "false", "Set to true to bypass initial pipeline validation upon creation. The pipeline still needs to be validated before activation. Defaults to false.")

	var createPipelineDescription string
	createPipelineCmd.Flags().StringVar(&createPipelineDescription, "description", "", "The description of the pipeline. Defaults to null.")

	var createPipelineInputDatafile string
	createPipelineCmd.Flags().StringVar(&createPipelineInputDatafile, "input-datafile", "", "This is a required parameter. The input data file that represents the pipeline.")

	streamsCmd.AddCommand(createTemplateCmd)

	var createTemplateDescription string
	createTemplateCmd.Flags().StringVar(&createTemplateDescription, "description", "", "This is a required parameter. Template description")
	createTemplateCmd.MarkFlagRequired("description")

	var createTemplateName string
	createTemplateCmd.Flags().StringVar(&createTemplateName, "name", "", "This is a required parameter. Template name")
	createTemplateCmd.MarkFlagRequired("name")

	var createTemplateInputDatafile string
	createTemplateCmd.Flags().StringVar(&createTemplateInputDatafile, "input-datafile", "", "This is a required parameter. The input data file that represents the pipeline.")

	streamsCmd.AddCommand(deactivatePipelineCmd)

	var deactivatePipelineId string
	deactivatePipelineCmd.Flags().StringVar(&deactivatePipelineId, "id", "", "This is a required parameter. Pipeline ID")
	deactivatePipelineCmd.MarkFlagRequired("id")

	var deactivatePipelineSkipSavepoint string
	deactivatePipelineCmd.Flags().StringVar(&deactivatePipelineSkipSavepoint, "skip-savepoint", "false", "Set to true to skip saving the state of a deactivated pipeline. When the pipeline is later activated, it will start with the newest data and skip any data that arrived after this deactivation, which can cause data loss. Defaults to false.")

	streamsCmd.AddCommand(decompileCmd)

	var decompileUpl string
	decompileCmd.Flags().StringVar(&decompileUpl, "upl", "", "This is a required parameter. ")
	decompileCmd.MarkFlagRequired("upl")

	streamsCmd.AddCommand(deleteConnectionCmd)

	var deleteConnectionConnectionId string
	deleteConnectionCmd.Flags().StringVar(&deleteConnectionConnectionId, "connection-id", "", "This is a required parameter. Connection ID")
	deleteConnectionCmd.MarkFlagRequired("connection-id")

	streamsCmd.AddCommand(deleteFileCmd)

	var deleteFileFileId string
	deleteFileCmd.Flags().StringVar(&deleteFileFileId, "file-id", "", "This is a required parameter. File ID")
	deleteFileCmd.MarkFlagRequired("file-id")

	streamsCmd.AddCommand(deleteLookupFileCmd)

	var deleteLookupFileFileId string
	deleteLookupFileCmd.Flags().StringVar(&deleteLookupFileFileId, "file-id", "", "This is a required parameter. File ID")
	deleteLookupFileCmd.MarkFlagRequired("file-id")

	streamsCmd.AddCommand(deletePipelineCmd)

	var deletePipelineId string
	deletePipelineCmd.Flags().StringVar(&deletePipelineId, "id", "", "This is a required parameter. Pipeline ID")
	deletePipelineCmd.MarkFlagRequired("id")

	streamsCmd.AddCommand(deleteTemplateCmd)

	var deleteTemplateTemplateId string
	deleteTemplateCmd.Flags().StringVar(&deleteTemplateTemplateId, "template-id", "", "This is a required parameter. Template ID")
	deleteTemplateCmd.MarkFlagRequired("template-id")

	streamsCmd.AddCommand(getFileMetadataCmd)

	var getFileMetadataFileId string
	getFileMetadataCmd.Flags().StringVar(&getFileMetadataFileId, "file-id", "", "This is a required parameter. File ID")
	getFileMetadataCmd.MarkFlagRequired("file-id")

	streamsCmd.AddCommand(getFilesMetadataCmd)

	streamsCmd.AddCommand(getInputSchemaCmd)

	var getInputSchemaNodeUuid string
	getInputSchemaCmd.Flags().StringVar(&getInputSchemaNodeUuid, "node-uuid", "", "This is a required parameter. The function ID.")
	getInputSchemaCmd.MarkFlagRequired("node-uuid")

	var getInputSchemaTargetPortName string
	getInputSchemaCmd.Flags().StringVar(&getInputSchemaTargetPortName, "target-port-name", "", "This is a required parameter. The name of the input port.")
	getInputSchemaCmd.MarkFlagRequired("target-port-name")

	var getInputSchemaInputDatafile string
	getInputSchemaCmd.Flags().StringVar(&getInputSchemaInputDatafile, "input-datafile", "", "This is a required parameter. The input data file that represents the pipeline.")

	streamsCmd.AddCommand(getLookupFileMetadataCmd)

	var getLookupFileMetadataFileId string
	getLookupFileMetadataCmd.Flags().StringVar(&getLookupFileMetadataFileId, "file-id", "", "This is a required parameter. File ID")
	getLookupFileMetadataCmd.MarkFlagRequired("file-id")

	streamsCmd.AddCommand(getLookupFilesMetadataCmd)

	streamsCmd.AddCommand(getLookupTableCmd)

	var getLookupTableConnectionId string
	getLookupTableCmd.Flags().StringVar(&getLookupTableConnectionId, "connection-id", "", "This is a required parameter. Connection ID")
	getLookupTableCmd.MarkFlagRequired("connection-id")

	var getLookupTableOffset int32
	getLookupTableCmd.Flags().Int32Var(&getLookupTableOffset, "offset", 0, "This is a required parameter. offset")
	getLookupTableCmd.MarkFlagRequired("offset")

	var getLookupTableSize int32
	getLookupTableCmd.Flags().Int32Var(&getLookupTableSize, "size", 0, "This is a required parameter. size")
	getLookupTableCmd.MarkFlagRequired("size")

	streamsCmd.AddCommand(getOutputSchemaCmd)

	var getOutputSchemaInputDatafile string
	getOutputSchemaCmd.Flags().StringVar(&getOutputSchemaInputDatafile, "input-datafile", "", "This is a required parameter. The input data file that represents the pipeline.")

	var getOutputSchemaNodeUuid string
	getOutputSchemaCmd.Flags().StringVar(&getOutputSchemaNodeUuid, "node-uuid", "", "The function ID. If omitted, returns the output schema for all functions.")

	var getOutputSchemaSourcePortName string
	getOutputSchemaCmd.Flags().StringVar(&getOutputSchemaSourcePortName, "source-port-name", "", "The name of the output port. Deprecated.")

	streamsCmd.AddCommand(getPipelineCmd)

	var getPipelineId string
	getPipelineCmd.Flags().StringVar(&getPipelineId, "id", "", "This is a required parameter. Pipeline ID")
	getPipelineCmd.MarkFlagRequired("id")

	var getPipelineVersion string
	getPipelineCmd.Flags().StringVar(&getPipelineVersion, "version", "", "version")

	streamsCmd.AddCommand(getPipelineLatestMetricsCmd)

	var getPipelineLatestMetricsId string
	getPipelineLatestMetricsCmd.Flags().StringVar(&getPipelineLatestMetricsId, "id", "", "This is a required parameter. Pipeline ID")
	getPipelineLatestMetricsCmd.MarkFlagRequired("id")

	streamsCmd.AddCommand(getPipelinesStatusCmd)

	var getPipelinesStatusActivated string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusActivated, "activated", "false", "activated")

	var getPipelinesStatusCreateUserId string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusCreateUserId, "create-user-id", "", "createUserId")

	var getPipelinesStatusName string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusName, "name", "", "name")

	var getPipelinesStatusOffset int32
	getPipelinesStatusCmd.Flags().Int32Var(&getPipelinesStatusOffset, "offset", 0, "offset")

	var getPipelinesStatusPageSize int32
	getPipelinesStatusCmd.Flags().Int32Var(&getPipelinesStatusPageSize, "page-size", 0, "pageSize")

	var getPipelinesStatusSortDir string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusSortDir, "sort-dir", "", "sortDir")

	var getPipelinesStatusSortField string
	getPipelinesStatusCmd.Flags().StringVar(&getPipelinesStatusSortField, "sort-field", "", "sortField")

	streamsCmd.AddCommand(getPreviewDataCmd)

	var getPreviewDataPreviewSessionId int64
	getPreviewDataCmd.Flags().Int64Var(&getPreviewDataPreviewSessionId, "preview-session-id", 0, "This is a required parameter. Preview Session ID")
	getPreviewDataCmd.MarkFlagRequired("preview-session-id")

	streamsCmd.AddCommand(getPreviewSessionCmd)

	var getPreviewSessionPreviewSessionId int64
	getPreviewSessionCmd.Flags().Int64Var(&getPreviewSessionPreviewSessionId, "preview-session-id", 0, "This is a required parameter. Preview Session ID")
	getPreviewSessionCmd.MarkFlagRequired("preview-session-id")

	streamsCmd.AddCommand(getPreviewSessionLatestMetricsCmd)

	var getPreviewSessionLatestMetricsPreviewSessionId int64
	getPreviewSessionLatestMetricsCmd.Flags().Int64Var(&getPreviewSessionLatestMetricsPreviewSessionId, "preview-session-id", 0, "This is a required parameter. Preview Session ID")
	getPreviewSessionLatestMetricsCmd.MarkFlagRequired("preview-session-id")

	streamsCmd.AddCommand(getRegistryCmd)

	var getRegistryLocal string
	getRegistryCmd.Flags().StringVar(&getRegistryLocal, "local", "true", "local")

	streamsCmd.AddCommand(getTemplateCmd)

	var getTemplateTemplateId string
	getTemplateCmd.Flags().StringVar(&getTemplateTemplateId, "template-id", "", "This is a required parameter. Template ID")
	getTemplateCmd.MarkFlagRequired("template-id")

	var getTemplateVersion int64
	getTemplateCmd.Flags().Int64Var(&getTemplateVersion, "version", 0, "Template version")

	streamsCmd.AddCommand(listConnectionsCmd)

	var listConnectionsConnectorId []string
	listConnectionsCmd.Flags().StringSliceVar(&listConnectionsConnectorId, "connector-id", nil, "")

	var listConnectionsCreateUserId string
	listConnectionsCmd.Flags().StringVar(&listConnectionsCreateUserId, "create-user-id", "", "")

	var listConnectionsFunctionId string
	listConnectionsCmd.Flags().StringVar(&listConnectionsFunctionId, "function-id", "", "")

	var listConnectionsFunctionOp string
	listConnectionsCmd.Flags().StringVar(&listConnectionsFunctionOp, "function-op", "", "")

	var listConnectionsName string
	listConnectionsCmd.Flags().StringVar(&listConnectionsName, "name", "", "")

	var listConnectionsOffset int32
	listConnectionsCmd.Flags().Int32Var(&listConnectionsOffset, "offset", 0, "")

	var listConnectionsPageSize int32
	listConnectionsCmd.Flags().Int32Var(&listConnectionsPageSize, "page-size", 0, "")

	var listConnectionsShowSecretNames string
	listConnectionsCmd.Flags().StringVar(&listConnectionsShowSecretNames, "show-secret-names", "", "")

	var listConnectionsSortDir string
	listConnectionsCmd.Flags().StringVar(&listConnectionsSortDir, "sort-dir", "", "Specify either ascending ('asc') or descending ('desc') sort order for a given field (sortField), which must be set for sortDir to apply. Defaults to 'asc'.")

	var listConnectionsSortField string
	listConnectionsCmd.Flags().StringVar(&listConnectionsSortField, "sort-field", "", "")

	streamsCmd.AddCommand(listConnectorsCmd)

	streamsCmd.AddCommand(listPipelinesCmd)

	var listPipelinesActivated string
	listPipelinesCmd.Flags().StringVar(&listPipelinesActivated, "activated", "false", "activated")

	var listPipelinesCreateUserId string
	listPipelinesCmd.Flags().StringVar(&listPipelinesCreateUserId, "create-user-id", "", "createUserId")

	var listPipelinesIncludeData string
	listPipelinesCmd.Flags().StringVar(&listPipelinesIncludeData, "include-data", "false", "includeData")

	var listPipelinesName string
	listPipelinesCmd.Flags().StringVar(&listPipelinesName, "name", "", "name")

	var listPipelinesOffset int32
	listPipelinesCmd.Flags().Int32Var(&listPipelinesOffset, "offset", 0, "offset")

	var listPipelinesPageSize int32
	listPipelinesCmd.Flags().Int32Var(&listPipelinesPageSize, "page-size", 0, "pageSize")

	var listPipelinesSortDir string
	listPipelinesCmd.Flags().StringVar(&listPipelinesSortDir, "sort-dir", "", "sortDir")

	var listPipelinesSortField string
	listPipelinesCmd.Flags().StringVar(&listPipelinesSortField, "sort-field", "", "sortField")

	streamsCmd.AddCommand(listTemplatesCmd)

	var listTemplatesOffset int32
	listTemplatesCmd.Flags().Int32Var(&listTemplatesOffset, "offset", 0, "offset")

	var listTemplatesPageSize int32
	listTemplatesCmd.Flags().Int32Var(&listTemplatesPageSize, "page-size", 0, "pageSize")

	var listTemplatesSortDir string
	listTemplatesCmd.Flags().StringVar(&listTemplatesSortDir, "sort-dir", "", "sortDir")

	var listTemplatesSortField string
	listTemplatesCmd.Flags().StringVar(&listTemplatesSortField, "sort-field", "", "sortField")

	streamsCmd.AddCommand(patchPipelineCmd)

	var patchPipelineId string
	patchPipelineCmd.Flags().StringVar(&patchPipelineId, "id", "", "This is a required parameter. Pipeline ID")
	patchPipelineCmd.MarkFlagRequired("id")

	var patchPipelineBypassValidation string
	patchPipelineCmd.Flags().StringVar(&patchPipelineBypassValidation, "bypass-validation", "false", "Set to true to bypass initial pipeline validation upon creation. The pipeline still needs to be validated before activation. Defaults to false.")

	var patchPipelineCreateUserId string
	patchPipelineCmd.Flags().StringVar(&patchPipelineCreateUserId, "create-user-id", "", "The user that created the pipeline. Deprecated.")

	var patchPipelineDescription string
	patchPipelineCmd.Flags().StringVar(&patchPipelineDescription, "description", "", "The description of the pipeline. Defaults to null.")

	var patchPipelineInputDatafile string
	patchPipelineCmd.Flags().StringVar(&patchPipelineInputDatafile, "input-datafile", "", "The input data file that represents the pipeline.")

	var patchPipelineName string
	patchPipelineCmd.Flags().StringVar(&patchPipelineName, "name", "", "The name of the pipeline.")

	streamsCmd.AddCommand(putConnectionCmd)

	var putConnectionConnectionId string
	putConnectionCmd.Flags().StringVar(&putConnectionConnectionId, "connection-id", "", "This is a required parameter. Connection ID")
	putConnectionCmd.MarkFlagRequired("connection-id")

	var putConnectionData string
	putConnectionCmd.Flags().StringVar(&putConnectionData, "data", "", "This is a required parameter. The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.")
	putConnectionCmd.MarkFlagRequired("data")

	var putConnectionDescription string
	putConnectionCmd.Flags().StringVar(&putConnectionDescription, "description", "", "This is a required parameter. The description of the connection.")
	putConnectionCmd.MarkFlagRequired("description")

	var putConnectionName string
	putConnectionCmd.Flags().StringVar(&putConnectionName, "name", "", "This is a required parameter. The name of the connection.")
	putConnectionCmd.MarkFlagRequired("name")

	streamsCmd.AddCommand(putTemplateCmd)

	var putTemplateData string
	putTemplateCmd.Flags().StringVar(&putTemplateData, "data", "", "This is a required parameter. ")
	putTemplateCmd.MarkFlagRequired("data")

	var putTemplateDescription string
	putTemplateCmd.Flags().StringVar(&putTemplateDescription, "description", "", "This is a required parameter. Template description")
	putTemplateCmd.MarkFlagRequired("description")

	var putTemplateName string
	putTemplateCmd.Flags().StringVar(&putTemplateName, "name", "", "This is a required parameter. Template name")
	putTemplateCmd.MarkFlagRequired("name")

	var putTemplateTemplateId string
	putTemplateCmd.Flags().StringVar(&putTemplateTemplateId, "template-id", "", "This is a required parameter. Template ID")
	putTemplateCmd.MarkFlagRequired("template-id")

	var putTemplateInputDatafile string
	putTemplateCmd.Flags().StringVar(&putTemplateInputDatafile, "input-datafile", "", "The input data file.")

	streamsCmd.AddCommand(reactivatePipelineCmd)

	var reactivatePipelineId string
	reactivatePipelineCmd.Flags().StringVar(&reactivatePipelineId, "id", "", "This is a required parameter. Pipeline ID")
	reactivatePipelineCmd.MarkFlagRequired("id")

	var reactivatePipelineActivateLatestVersion string
	reactivatePipelineCmd.Flags().StringVar(&reactivatePipelineActivateLatestVersion, "activate-latest-version", "false", "Set to true to activate the latest version of the pipeline. Set to false to use the previously activated version of the pipeline. Defaults to true.")

	var reactivatePipelineAllowNonRestoredState string
	reactivatePipelineCmd.Flags().StringVar(&reactivatePipelineAllowNonRestoredState, "allow-non-restored-state", "false", "Set to true to allow the pipeline to ignore any unused progress states. In some cases, when a data pipeline is changed, the progress state will be stored for functions that no longer exist, so this must be set to reactivate a pipeline in this state. Defaults to false.")

	var reactivatePipelineSkipRestoreState string
	reactivatePipelineCmd.Flags().StringVar(&reactivatePipelineSkipRestoreState, "skip-restore-state", "false", "Set to true to start reading from the latest input rather than from where the pipeline's previous run left off, which can cause data loss. Defaults to false.")

	streamsCmd.AddCommand(startPreviewCmd)

	var startPreviewInputDatafile string
	startPreviewCmd.Flags().StringVar(&startPreviewInputDatafile, "input-datafile", "", "The input data file that represents the pipeline.")

	var startPreviewRecordsLimit int32
	startPreviewCmd.Flags().Int32Var(&startPreviewRecordsLimit, "records-limit", 0, "The maximum number of events per function. Defaults to 100.")

	var startPreviewRecordsPerPipeline int32
	startPreviewCmd.Flags().Int32Var(&startPreviewRecordsPerPipeline, "records-per-pipeline", 0, "The maximum number of events per pipeline. Defaults to 10000.")

	var startPreviewSessionLifetimeMs int64
	startPreviewCmd.Flags().Int64Var(&startPreviewSessionLifetimeMs, "session-lifetime-ms", 0, "The maximum lifetime of a session, in milliseconds. Defaults to 300,000.")

	streamsCmd.AddCommand(stopPreviewCmd)

	var stopPreviewPreviewSessionId int64
	stopPreviewCmd.Flags().Int64Var(&stopPreviewPreviewSessionId, "preview-session-id", 0, "This is a required parameter. Preview Session ID")
	stopPreviewCmd.MarkFlagRequired("preview-session-id")

	streamsCmd.AddCommand(updateConnectionCmd)

	var updateConnectionConnectionId string
	updateConnectionCmd.Flags().StringVar(&updateConnectionConnectionId, "connection-id", "", "This is a required parameter. Connection ID")
	updateConnectionCmd.MarkFlagRequired("connection-id")

	var updateConnectionData string
	updateConnectionCmd.Flags().StringVar(&updateConnectionData, "data", "", "The key-value pairs of configurations for this connection. Connectors may have some configurations that are required, which all connections must provide values for. For configuration values of type BYTES, the provided values must be Base64 encoded.")

	var updateConnectionDescription string
	updateConnectionCmd.Flags().StringVar(&updateConnectionDescription, "description", "", "The description of the connection.")

	var updateConnectionName string
	updateConnectionCmd.Flags().StringVar(&updateConnectionName, "name", "", "The name of the connection.")

	streamsCmd.AddCommand(updatePipelineCmd)

	var updatePipelineId string
	updatePipelineCmd.Flags().StringVar(&updatePipelineId, "id", "", "This is a required parameter. Pipeline ID")
	updatePipelineCmd.MarkFlagRequired("id")

	var updatePipelineName string
	updatePipelineCmd.Flags().StringVar(&updatePipelineName, "name", "", "This is a required parameter. The name of the pipeline.")
	updatePipelineCmd.MarkFlagRequired("name")

	var updatePipelineBypassValidation string
	updatePipelineCmd.Flags().StringVar(&updatePipelineBypassValidation, "bypass-validation", "false", "Set to true to bypass initial pipeline validation upon creation. The pipeline still needs to be validated before activation. Defaults to false.")

	var updatePipelineDescription string
	updatePipelineCmd.Flags().StringVar(&updatePipelineDescription, "description", "", "The description of the pipeline. Defaults to null.")

	var updatePipelineInputDatafile string
	updatePipelineCmd.Flags().StringVar(&updatePipelineInputDatafile, "input-datafile", "", "This is a required parameter. The input data file that represents the pipeline.")

	streamsCmd.AddCommand(updateTemplateCmd)

	var updateTemplateTemplateId string
	updateTemplateCmd.Flags().StringVar(&updateTemplateTemplateId, "template-id", "", "This is a required parameter. Template ID")
	updateTemplateCmd.MarkFlagRequired("template-id")

	var updateTemplateDescription string
	updateTemplateCmd.Flags().StringVar(&updateTemplateDescription, "description", "", "Template description")

	var updateTemplateInputDatafile string
	updateTemplateCmd.Flags().StringVar(&updateTemplateInputDatafile, "input-datafile", "", "The input data file.")

	var updateTemplateName string
	updateTemplateCmd.Flags().StringVar(&updateTemplateName, "name", "", "Template name")

	streamsCmd.AddCommand(uploadFileCmd)

	var uploadFileFileName string
	uploadFileCmd.Flags().StringVar(&uploadFileFileName, "file-name", "", "File to upload.")

	streamsCmd.AddCommand(uploadLookupFileCmd)

	streamsCmd.AddCommand(validatePipelineCmd)

	var validatePipelineInputDatafile string
	validatePipelineCmd.Flags().StringVar(&validatePipelineInputDatafile, "input-datafile", "", "The input data file that represents the pipeline.")

}
