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
 * Identity
 *
 * With the Identity service in Splunk Cloud Services, you can authenticate and authorize Splunk Cloud Services users.
 *
 * API version: v3.1 (recommended default)
 * Generated by: OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.
 */

package identity

type AddGroupMemberBody struct {
	Name string `json:"name"`
}

type AddGroupRoleBody struct {
	Name string `json:"name"`
}

type AddMemberBody struct {
	Name string `json:"name"`
}

type AddRolePermissionBody struct {
	Permission string `json:"permission"`
}

type CreateGroupBody struct {
	Name string `json:"name"`
}

// Payload when creating a principal
type CreatePrincipalBody struct {
	Kind        PrincipalKind           `json:"kind"`
	AcceptTos   *bool                   `json:"acceptTos,omitempty"`
	Credentials *CredentialList         `json:"credentials,omitempty"`
	Enabled     *bool                   `json:"enabled,omitempty"`
	Key         *EcJwk                  `json:"key,omitempty"`
	Name        *string                 `json:"name,omitempty"`
	Profile     *CreatePrincipalProfile `json:"profile,omitempty"`
}

// Payload when creating Principal profile
type CreatePrincipalProfile struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type CreateRoleBody struct {
	Name string `json:"name"`
}

// Principal credential
type Credential struct {
	Type  CredentialType `json:"type"`
	Value string         `json:"value"`
}

type CredentialType string

// List of CredentialType
const (
	CredentialTypePassword CredentialType = "password"
)

// List of credentials
type CredentialList []Credential

type EcJwk struct {
	Alg *EcJwkAlg `json:"alg,omitempty"`
	Crv *string   `json:"crv,omitempty"`
	D   *string   `json:"d,omitempty"`
	Kid *string   `json:"kid,omitempty"`
	Kty *EcJwkKty `json:"kty,omitempty"`
	X   *string   `json:"x,omitempty"`
	Y   *string   `json:"y,omitempty"`
}

type EcJwkAlg string

// List of EcJwkAlg
const (
	EcJwkAlgEs256 EcJwkAlg = "ES256"
	EcJwkAlgEs384 EcJwkAlg = "ES384"
	EcJwkAlgEs512 EcJwkAlg = "ES512"
)

type EcJwkKty string

// List of EcJwkKty
const (
	EcJwkKtyEc EcJwkKty = "EC"
)

type Group struct {
	Name        string  `json:"name"`
	CreatedAt   *string `json:"createdAt,omitempty"`
	CreatedBy   *string `json:"createdBy,omitempty"`
	MemberCount *int32  `json:"memberCount,omitempty"`
	RoleCount   *int32  `json:"roleCount,omitempty"`
	Tenant      *string `json:"tenant,omitempty"`
}

// A list of group.
type GroupList struct {
	Items    []Group `json:"items"`
	NextLink string  `json:"nextLink"`
}

// Represents a member that belongs to a group
type GroupMember struct {
	AddedAt   string `json:"addedAt"`
	AddedBy   string `json:"addedBy"`
	Group     string `json:"group"`
	Principal string `json:"principal"`
	Tenant    string `json:"tenant"`
}

// A list of members belonging to a particular group.
type GroupMemberList struct {
	Items    []GroupMember `json:"items"`
	NextLink string        `json:"nextLink"`
}

// Represents a role that is assigned to a group
type GroupRole struct {
	AddedAt string `json:"addedAt"`
	AddedBy string `json:"addedBy"`
	Group   string `json:"group"`
	Role    string `json:"role"`
	Tenant  string `json:"tenant"`
}

// A list of group role.
type GroupRoleList struct {
	Items    []GroupRole `json:"items"`
	NextLink string      `json:"nextLink"`
}

type IdentityProviderBody struct {
	Config      IdentityProviderBodyConfig `json:"config"`
	Id          string                     `json:"id"`
	Kind        IdentityProviderBodyKind   `json:"kind"`
	Title       string                     `json:"title"`
	CreatedAt   *string                    `json:"createdAt,omitempty"`
	CreatedBy   *string                    `json:"createdBy,omitempty"`
	Description *string                    `json:"description,omitempty"`
	Enabled     *bool                      `json:"enabled,omitempty"`
}

type IdentityProviderBodyKind string

// List of IdentityProviderBodyKind
const (
	IdentityProviderBodyKindSaml IdentityProviderBodyKind = "saml"
)

type IdentityProviderBodyConfig struct {
	Certificate            string                           `json:"certificate"`
	EmailAttribute         string                           `json:"email_attribute"`
	EntityDescriptor       string                           `json:"entity_descriptor"`
	Method                 IdentityProviderBodyConfigMethod `json:"method"`
	SingleSignOnServiceUrl string                           `json:"single_sign_on_service_url"`
	FirstNameAttribute     *string                          `json:"first_name_attribute,omitempty"`
	LastNameAttribute      *string                          `json:"last_name_attribute,omitempty"`
}

type IdentityProviderBodyConfigMethod string

// List of IdentityProviderBodyConfigMethod
const (
	IdentityProviderBodyConfigMethodPost     IdentityProviderBodyConfigMethod = "post"
	IdentityProviderBodyConfigMethodRedirect IdentityProviderBodyConfigMethod = "redirect"
)

type IdentityProviderConfigBody struct {
	Config      IdentityProviderConfigBodyConfig `json:"config"`
	Id          string                           `json:"id"`
	Kind        IdentityProviderConfigBodyKind   `json:"kind"`
	Description *string                          `json:"description,omitempty"`
	Enabled     *bool                            `json:"enabled,omitempty"`
	Title       *string                          `json:"title,omitempty"`
}

type IdentityProviderConfigBodyKind string

// List of IdentityProviderConfigBodyKind
const (
	IdentityProviderConfigBodyKindSaml IdentityProviderConfigBodyKind = "saml"
)

type IdentityProviderConfigBodyConfig struct {
	Certificate            *string                                 `json:"certificate,omitempty"`
	EmailAttribute         *string                                 `json:"email_attribute,omitempty"`
	EntityDescriptor       *string                                 `json:"entity_descriptor,omitempty"`
	FirstNameAttribute     *string                                 `json:"first_name_attribute,omitempty"`
	LastNameAttribute      *string                                 `json:"last_name_attribute,omitempty"`
	Method                 *IdentityProviderConfigBodyConfigMethod `json:"method,omitempty"`
	SingleSignOnServiceUrl *string                                 `json:"single_sign_on_service_url,omitempty"`
}

type IdentityProviderConfigBodyConfigMethod string

// List of IdentityProviderConfigBodyConfigMethod
const (
	IdentityProviderConfigBodyConfigMethodPost     IdentityProviderConfigBodyConfigMethod = "post"
	IdentityProviderConfigBodyConfigMethodRedirect IdentityProviderConfigBodyConfigMethod = "redirect"
)

// Represents a member that belongs to a tenant.
type Member struct {
	Name string `json:"name"`
	// When the principal was added to the tenant.
	AddedAt    *string           `json:"addedAt,omitempty"`
	AddedBy    *string           `json:"addedBy,omitempty"`
	ExpiresAt  *string           `json:"expiresAt,omitempty"`
	GroupCount *int32            `json:"groupCount,omitempty"`
	Profile    *PrincipalProfile `json:"profile,omitempty"`
	Tenant     *string           `json:"tenant,omitempty"`
	Visible    *bool             `json:"visible,omitempty"`
}

// A list of Members belonging to a particular Tenant.
type MemberList struct {
	Items    []Member `json:"items"`
	NextLink string   `json:"nextLink"`
}

// A list of permissions.
type PermissionList struct {
	Items    []string `json:"items"`
	NextLink string   `json:"nextLink"`
}

type PermissionString string

type Principal struct {
	Name      string            `json:"name"`
	CreatedAt *string           `json:"createdAt,omitempty"`
	CreatedBy *string           `json:"createdBy,omitempty"`
	Kind      *PrincipalKind    `json:"kind,omitempty"`
	Profile   *PrincipalProfile `json:"profile,omitempty"`
	Tenants   []string          `json:"tenants,omitempty"`
	UpdatedAt *string           `json:"updatedAt,omitempty"`
	UpdatedBy *string           `json:"updatedBy,omitempty"`
}

// PrincipalKind :
type PrincipalKind string

// List of PrincipalKind
const (
	PrincipalKindUser           PrincipalKind = "user"
	PrincipalKindServiceAccount PrincipalKind = "service_account"
	PrincipalKindService        PrincipalKind = "service"
)

// A list of Principals.
type PrincipalList struct {
	Items    []Principal `json:"items"`
	NextLink string      `json:"nextLink"`
}

// Profile information for a principal
type PrincipalProfile struct {
	Email     *string `json:"email,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	FullName  *string `json:"fullName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

type PrincipalPublicKey struct {
	CreatedAt string                   `json:"createdAt"`
	CreatedBy string                   `json:"createdBy"`
	Key       EcJwk                    `json:"key"`
	Status    PrincipalPublicKeyStatus `json:"status"`
	UpdatedAt string                   `json:"updatedAt"`
	UpdatedBy string                   `json:"updatedBy"`
}

type PrincipalPublicKeyStatus string

// List of PrincipalPublicKeyStatus
const (
	PrincipalPublicKeyStatusActive   PrincipalPublicKeyStatus = "active"
	PrincipalPublicKeyStatusInactive PrincipalPublicKeyStatus = "inactive"
	PrincipalPublicKeyStatusDeleted  PrincipalPublicKeyStatus = "deleted"
)

type PrincipalPublicKeyStatusBody struct {
	Status PrincipalPublicKeyStatusBodyStatus `json:"status"`
}

type PrincipalPublicKeyStatusBodyStatus string

// List of PrincipalPublicKeyStatusBodyStatus
const (
	PrincipalPublicKeyStatusBodyStatusActive   PrincipalPublicKeyStatusBodyStatus = "active"
	PrincipalPublicKeyStatusBodyStatusInactive PrincipalPublicKeyStatusBodyStatus = "inactive"
)

type PrincipalPublicKeys []PrincipalPublicKey

type Role struct {
	Name        string   `json:"name"`
	CreatedAt   *string  `json:"createdAt,omitempty"`
	CreatedBy   *string  `json:"createdBy,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
	Tenant      *string  `json:"tenant,omitempty"`
}

// A list of role.
type RoleList struct {
	Items    []Role `json:"items"`
	NextLink string `json:"nextLink"`
}

type RolePermission struct {
	Permission string  `json:"permission"`
	AddedAt    *string `json:"addedAt,omitempty"`
	AddedBy    *string `json:"addedBy,omitempty"`
	Role       *string `json:"role,omitempty"`
	Tenant     *string `json:"tenant,omitempty"`
}

// A list of role permissions.
type RolePermissionList struct {
	Items    []RolePermission `json:"items"`
	NextLink string           `json:"nextLink"`
}

type Tenant struct {
	Name      string        `json:"name"`
	CreatedAt *string       `json:"createdAt,omitempty"`
	CreatedBy *string       `json:"createdBy,omitempty"`
	Status    *TenantStatus `json:"status,omitempty"`
}

type TenantName string

// TenantStatus :
type TenantStatus string

// List of TenantStatus
const (
	TenantStatusProvisioning TenantStatus = "provisioning"
	TenantStatusFailed       TenantStatus = "failed"
	TenantStatusReady        TenantStatus = "ready"
	TenantStatusDeleting     TenantStatus = "deleting"
	TenantStatusDeleted      TenantStatus = "deleted"
	TenantStatusTombstoned   TenantStatus = "tombstoned"
	TenantStatusSuspended    TenantStatus = "suspended"
)

type ValidateInfo struct {
	ClientId  string           `json:"clientId"`
	Kind      ValidateInfoKind `json:"kind"`
	Name      string           `json:"name"`
	Principal *Principal       `json:"principal,omitempty"`
	Tenant    *Tenant          `json:"tenant,omitempty"`
}

type ValidateInfoKind string

// List of ValidateInfoKind
const (
	ValidateInfoKindPrincipal ValidateInfoKind = "principal"
	ValidateInfoKindApiKey    ValidateInfoKind = "api_key"
)
