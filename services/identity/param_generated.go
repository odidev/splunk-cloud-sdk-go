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
 *
 * Identity
 *
 * With the Identity service in Splunk Cloud Services, you can authenticate and authorize Splunk Cloud Services users.
 *
 * API version: v2beta1.13 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package identity

// ValidateTokenQueryParams represents valid query parameters for the ValidateToken operation
// For convenience ValidateTokenQueryParams can be formed in a single statement, for example:
//     `v := ValidateTokenQueryParams{}.SetInclude(...)`
type ValidateTokenQueryParams struct {
	// Include : Include additional information to return when validating tenant membership. Valid parameters [tenant, principal]
	Include ValidateTokeninclude `key:"include"`
}

func (q ValidateTokenQueryParams) SetInclude(v ValidateTokeninclude) ValidateTokenQueryParams {
	q.Include = v
	return q
}

// ValidateTokeninclude : Include additional information to return when validating tenant membership. Valid parameters [tenant, principal]
type ValidateTokeninclude []string

// List of ValidateTokeninclude values
const (
	ValidateTokenincludeTenant    ValidateTokeninclude = "tenant"
	ValidateTokenincludePrincipal ValidateTokeninclude = "principal"
)
