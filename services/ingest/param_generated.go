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
 * API version: v1beta2.32 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package ingest

// ListCollectorTokensQueryParams represents valid query parameters for the ListCollectorTokens operation
// For convenience ListCollectorTokensQueryParams can be formed in a single statement, for example:
//     `v := ListCollectorTokensQueryParams{}.SetLimit(...).SetOffset(...)`
type ListCollectorTokensQueryParams struct {
	Limit  *int64 `key:"limit"`
	Offset *int64 `key:"offset"`
}

func (q ListCollectorTokensQueryParams) SetLimit(v int64) ListCollectorTokensQueryParams {
	q.Limit = &v
	return q
}

func (q ListCollectorTokensQueryParams) SetOffset(v int64) ListCollectorTokensQueryParams {
	q.Offset = &v
	return q
}
