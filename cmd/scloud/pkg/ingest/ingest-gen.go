// Package ingest -- generated by scloudgen
// !! DO NOT EDIT !!
//
package ingest

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/auth"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/flags"
	"github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/jsonx"
	model "github.com/splunk/splunk-cloud-sdk-go/services/ingest"
)

// DeleteAllCollectorTokens Delete All dsphec tokens for a given tenant.
func DeleteAllCollectorTokens(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}

	resp, err := client.IngestService.DeleteAllCollectorTokens()
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// DeleteCollectorToken Delete dsphec token by name.
func DeleteCollectorToken(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var tokenName string
	err = flags.ParseFlag(cmd.Flags(), "token-name", &tokenName)
	if err != nil {
		return fmt.Errorf(`error parsing "token-name": ` + err.Error())
	}

	resp, err := client.IngestService.DeleteCollectorToken(tokenName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// GetCollectorToken Get the metadata of a dsphec token by name.
func GetCollectorToken(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var tokenName string
	err = flags.ParseFlag(cmd.Flags(), "token-name", &tokenName)
	if err != nil {
		return fmt.Errorf(`error parsing "token-name": ` + err.Error())
	}

	resp, err := client.IngestService.GetCollectorToken(tokenName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// ListCollectorTokens List dsphec tokens for a tenant.
func ListCollectorTokens(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var limitDefault int64
	limit := &limitDefault
	err = flags.ParseFlag(cmd.Flags(), "limit", &limit)
	if err != nil {
		return fmt.Errorf(`error parsing "limit": ` + err.Error())
	}
	var offsetDefault int64
	offset := &offsetDefault
	err = flags.ParseFlag(cmd.Flags(), "offset", &offset)
	if err != nil {
		return fmt.Errorf(`error parsing "offset": ` + err.Error())
	}
	// Form query params
	generated_query := model.ListCollectorTokensQueryParams{}
	generated_query.Limit = limit
	generated_query.Offset = offset

	resp, err := client.IngestService.ListCollectorTokens(&generated_query)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// PostCollectorTokens Creates dsphec tokens.
func PostCollectorTokens(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var descriptionDefault string
	description := &descriptionDefault
	err = flags.ParseFlag(cmd.Flags(), "description", &description)
	if err != nil {
		return fmt.Errorf(`error parsing "description": ` + err.Error())
	}
	var indexDefault string
	index := &indexDefault
	err = flags.ParseFlag(cmd.Flags(), "index", &index)
	if err != nil {
		return fmt.Errorf(`error parsing "index": ` + err.Error())
	}
	var name string
	err = flags.ParseFlag(cmd.Flags(), "name", &name)
	if err != nil {
		return fmt.Errorf(`error parsing "name": ` + err.Error())
	}
	var sourceDefault string
	source := &sourceDefault
	err = flags.ParseFlag(cmd.Flags(), "source", &source)
	if err != nil {
		return fmt.Errorf(`error parsing "source": ` + err.Error())
	}
	var sourcetypeDefault string
	sourcetype := &sourcetypeDefault
	err = flags.ParseFlag(cmd.Flags(), "sourcetype", &sourcetype)
	if err != nil {
		return fmt.Errorf(`error parsing "sourcetype": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.HecTokenCreateRequest{

		Description: description,
		Index:       index,
		Name:        name,
		Source:      source,
		Sourcetype:  sourcetype,
	}

	resp, err := client.IngestService.PostCollectorTokens(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// PostEvents Sends events.
func PostEvents(cmd *cobra.Command, args []string) error {

	var err error

	// Parse all flags

	var attributes map[string]interface{}
	err = flags.ParseFlag(cmd.Flags(), "attributes", &attributes)
	if err != nil {
		return fmt.Errorf(`error parsing "attributes": ` + err.Error())
	}
	var format string
	err = flags.ParseFlag(cmd.Flags(), "format", &format)
	if err != nil {
		return fmt.Errorf(`error parsing "format": ` + err.Error())
	}
	var hostDefault string
	host := &hostDefault
	err = flags.ParseFlag(cmd.Flags(), "host", &host)
	if err != nil {
		return fmt.Errorf(`error parsing "host": ` + err.Error())
	}
	var idDefault string
	id := &idDefault
	err = flags.ParseFlag(cmd.Flags(), "id", &id)
	if err != nil {
		return fmt.Errorf(`error parsing "id": ` + err.Error())
	}
	var nanosDefault int32
	nanos := &nanosDefault
	err = flags.ParseFlag(cmd.Flags(), "nanos", &nanos)
	if err != nil {
		return fmt.Errorf(`error parsing "nanos": ` + err.Error())
	}
	var sourceDefault string
	source := &sourceDefault
	err = flags.ParseFlag(cmd.Flags(), "source", &source)
	if err != nil {
		return fmt.Errorf(`error parsing "source": ` + err.Error())
	}
	var sourcetypeDefault string
	sourcetype := &sourcetypeDefault
	err = flags.ParseFlag(cmd.Flags(), "sourcetype", &sourcetype)
	if err != nil {
		return fmt.Errorf(`error parsing "sourcetype": ` + err.Error())
	}
	var timestampDefault int64
	timestamp := &timestampDefault
	err = flags.ParseFlag(cmd.Flags(), "timestamp", &timestamp)
	if err != nil {
		return fmt.Errorf(`error parsing "timestamp": ` + err.Error())
	}
	// Form the request body
	generated_request_body := []model.Event{
		{
			Attributes: attributes,

			Host:       host,
			Id:         id,
			Nanos:      nanos,
			Source:     source,
			Sourcetype: sourcetype,
			Timestamp:  timestamp,
		},
	}

	resp, err := PostEventsOverride(generated_request_body, format)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// PostMetrics Sends metric events.
func PostMetrics(cmd *cobra.Command, args []string) error {

	var err error

	// Parse all flags

	var defaultDimensions map[string]string
	err = flags.ParseFlag(cmd.Flags(), "default-dimensions", &defaultDimensions)
	if err != nil {
		return fmt.Errorf(`error parsing "default-dimensions": ` + err.Error())
	}
	var defaultTypeDefault string
	defaultType := &defaultTypeDefault
	err = flags.ParseFlag(cmd.Flags(), "default-type", &defaultType)
	if err != nil {
		return fmt.Errorf(`error parsing "default-type": ` + err.Error())
	}
	var defaultUnitDefault string
	defaultUnit := &defaultUnitDefault
	err = flags.ParseFlag(cmd.Flags(), "default-unit", &defaultUnit)
	if err != nil {
		return fmt.Errorf(`error parsing "default-unit": ` + err.Error())
	}
	var hostDefault string
	host := &hostDefault
	err = flags.ParseFlag(cmd.Flags(), "host", &host)
	if err != nil {
		return fmt.Errorf(`error parsing "host": ` + err.Error())
	}
	var idDefault string
	id := &idDefault
	err = flags.ParseFlag(cmd.Flags(), "id", &id)
	if err != nil {
		return fmt.Errorf(`error parsing "id": ` + err.Error())
	}
	var nanosDefault int32
	nanos := &nanosDefault
	err = flags.ParseFlag(cmd.Flags(), "nanos", &nanos)
	if err != nil {
		return fmt.Errorf(`error parsing "nanos": ` + err.Error())
	}
	var sourceDefault string
	source := &sourceDefault
	err = flags.ParseFlag(cmd.Flags(), "source", &source)
	if err != nil {
		return fmt.Errorf(`error parsing "source": ` + err.Error())
	}
	var sourcetypeDefault string
	sourcetype := &sourcetypeDefault
	err = flags.ParseFlag(cmd.Flags(), "sourcetype", &sourcetype)
	if err != nil {
		return fmt.Errorf(`error parsing "sourcetype": ` + err.Error())
	}
	var timestampDefault int64
	timestamp := &timestampDefault
	err = flags.ParseFlag(cmd.Flags(), "timestamp", &timestamp)
	if err != nil {
		return fmt.Errorf(`error parsing "timestamp": ` + err.Error())
	}
	// Form the request body
	generated_request_body := []model.MetricEvent{
		{
			Attributes: &model.MetricAttribute{
				DefaultDimensions: defaultDimensions,
				DefaultType:       defaultType,
				DefaultUnit:       defaultUnit,
			},

			Host:       host,
			Id:         id,
			Nanos:      nanos,
			Source:     source,
			Sourcetype: sourcetype,
			Timestamp:  timestamp,
		},
	}

	resp, err := PostMetricsOverride(generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// PutCollectorToken Update the metadata of a dsphec token by name.
func PutCollectorToken(cmd *cobra.Command, args []string) error {

	client, err := auth.GetClient()
	if err != nil {
		return err
	}
	// Parse all flags

	var descriptionDefault string
	description := &descriptionDefault
	err = flags.ParseFlag(cmd.Flags(), "description", &description)
	if err != nil {
		return fmt.Errorf(`error parsing "description": ` + err.Error())
	}
	var indexDefault string
	index := &indexDefault
	err = flags.ParseFlag(cmd.Flags(), "index", &index)
	if err != nil {
		return fmt.Errorf(`error parsing "index": ` + err.Error())
	}
	var sourceDefault string
	source := &sourceDefault
	err = flags.ParseFlag(cmd.Flags(), "source", &source)
	if err != nil {
		return fmt.Errorf(`error parsing "source": ` + err.Error())
	}
	var sourcetypeDefault string
	sourcetype := &sourcetypeDefault
	err = flags.ParseFlag(cmd.Flags(), "sourcetype", &sourcetype)
	if err != nil {
		return fmt.Errorf(`error parsing "sourcetype": ` + err.Error())
	}
	var tokenName string
	err = flags.ParseFlag(cmd.Flags(), "token-name", &tokenName)
	if err != nil {
		return fmt.Errorf(`error parsing "token-name": ` + err.Error())
	}
	// Form the request body
	generated_request_body := model.HecTokenUpdateRequest{

		Description: description,
		Index:       index,
		Source:      source,
		Sourcetype:  sourcetype,
	}

	resp, err := client.IngestService.PutCollectorToken(tokenName, generated_request_body)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}

// UploadFiles Upload a CSV or text file that contains events.
func UploadFiles(cmd *cobra.Command, args []string) error {

	var err error

	// Parse all flags

	var fileName string
	err = flags.ParseFlag(cmd.Flags(), "file-name", &fileName)
	if err != nil {
		return fmt.Errorf(`error parsing "file-name": ` + err.Error())
	}

	resp, err := UploadFilesOverride(fileName)
	if err != nil {
		return err
	}
	jsonx.Pprint(cmd, resp)
	return nil
}
