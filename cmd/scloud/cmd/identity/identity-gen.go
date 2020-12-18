// Package identity -- generated by scloudgen
// !! DO NOT EDIT !!
//
package identity

import (
	"github.com/spf13/cobra"
	impl "github.com/splunk/splunk-cloud-sdk-go/cmd/scloud/pkg/identity"
)

// addGroupMember -- Adds a member to a given group.
var addGroupMemberCmd = &cobra.Command{
	Use:   "add-group-member",
	Short: "Adds a member to a given group.",
	RunE:  impl.AddGroupMember,
}

// addGroupRole -- Adds a role to a given group.
var addGroupRoleCmd = &cobra.Command{
	Use:   "add-group-role",
	Short: "Adds a role to a given group.",
	RunE:  impl.AddGroupRole,
}

// addMember -- Adds a member to a given tenant.
var addMemberCmd = &cobra.Command{
	Use:   "add-member",
	Short: "Adds a member to a given tenant.",
	RunE:  impl.AddMember,
}

// addPrincipalPublicKey -- Add service principal public key
var addPrincipalPublicKeyCmd = &cobra.Command{
	Use:   "add-principal-public-key",
	Short: "Add service principal public key",
	RunE:  impl.AddPrincipalPublicKey,
}

// addRolePermission -- Adds permissions to a role in a given tenant.
var addRolePermissionCmd = &cobra.Command{
	Use:   "add-role-permission",
	Short: "Adds permissions to a role in a given tenant.",
	RunE:  impl.AddRolePermission,
}

// createGroup -- Creates a new group in a given tenant.
var createGroupCmd = &cobra.Command{
	Use:   "create-group",
	Short: "Creates a new group in a given tenant.",
	RunE:  impl.CreateGroup,
}

// createPrincipal -- Create a new principal
var createPrincipalCmd = &cobra.Command{
	Use:   "create-principal",
	Short: "Create a new principal",
	RunE:  impl.CreatePrincipal,
}

// createRole -- Creates a new authorization role in a given tenant.
var createRoleCmd = &cobra.Command{
	Use:   "create-role",
	Short: "Creates a new authorization role in a given tenant.",
	RunE:  impl.CreateRole,
}

// deleteGroup -- Deletes a group in a given tenant.
var deleteGroupCmd = &cobra.Command{
	Use:   "delete-group",
	Short: "Deletes a group in a given tenant.",
	RunE:  impl.DeleteGroup,
}

// deletePrincipalPublicKey -- Deletes principal public key
var deletePrincipalPublicKeyCmd = &cobra.Command{
	Use:   "delete-principal-public-key",
	Short: "Deletes principal public key",
	RunE:  impl.DeletePrincipalPublicKey,
}

// deleteRole -- Deletes a defined role for a given tenant.
var deleteRoleCmd = &cobra.Command{
	Use:   "delete-role",
	Short: "Deletes a defined role for a given tenant.",
	RunE:  impl.DeleteRole,
}

// getGroup -- Returns information about a given group within a tenant.
var getGroupCmd = &cobra.Command{
	Use:   "get-group",
	Short: "Returns information about a given group within a tenant.",
	RunE:  impl.GetGroup,
}

// getGroupMember -- Returns information about a given member within a given group.
var getGroupMemberCmd = &cobra.Command{
	Use:   "get-group-member",
	Short: "Returns information about a given member within a given group.",
	RunE:  impl.GetGroupMember,
}

// getGroupRole -- Returns information about a given role within a given group.
var getGroupRoleCmd = &cobra.Command{
	Use:   "get-group-role",
	Short: "Returns information about a given role within a given group.",
	RunE:  impl.GetGroupRole,
}

// getMember -- Returns a member of a given tenant.
var getMemberCmd = &cobra.Command{
	Use:   "get-member",
	Short: "Returns a member of a given tenant.",
	RunE:  impl.GetMember,
}

// getPrincipal -- Returns the details of a principal, including its tenant membership and any relevant profile information.

var getPrincipalCmd = &cobra.Command{
	Use:   "get-principal",
	Short: "Returns the details of a principal, including its tenant membership and any relevant profile information.",
	RunE:  impl.GetPrincipal,
}

// getPrincipalPublicKey -- Returns principal public key
var getPrincipalPublicKeyCmd = &cobra.Command{
	Use:   "get-principal-public-key",
	Short: "Returns principal public key",
	RunE:  impl.GetPrincipalPublicKey,
}

// getPrincipalPublicKeys -- Returns principal public keys
var getPrincipalPublicKeysCmd = &cobra.Command{
	Use:   "get-principal-public-keys",
	Short: "Returns principal public keys",
	RunE:  impl.GetPrincipalPublicKeys,
}

// getRole -- Returns a role for a given tenant.
var getRoleCmd = &cobra.Command{
	Use:   "get-role",
	Short: "Returns a role for a given tenant.",
	RunE:  impl.GetRole,
}

// getRolePermission -- Gets a permission for the specified role.
var getRolePermissionCmd = &cobra.Command{
	Use:   "get-role-permission",
	Short: "Gets a permission for the specified role.",
	RunE:  impl.GetRolePermission,
}

// listGroupMembers -- Returns a list of the members within a given group.
var listGroupMembersCmd = &cobra.Command{
	Use:   "list-group-members",
	Short: "Returns a list of the members within a given group.",
	RunE:  impl.ListGroupMembers,
}

// listGroupRoles -- Returns a list of the roles that are attached to a group within a given tenant.
var listGroupRolesCmd = &cobra.Command{
	Use:   "list-group-roles",
	Short: "Returns a list of the roles that are attached to a group within a given tenant.",
	RunE:  impl.ListGroupRoles,
}

// listGroups -- List the groups that exist in a given tenant.
var listGroupsCmd = &cobra.Command{
	Use:   "list-groups",
	Short: "List the groups that exist in a given tenant.",
	RunE:  impl.ListGroups,
}

// listMemberGroups -- Returns a list of groups that a member belongs to within a tenant.
var listMemberGroupsCmd = &cobra.Command{
	Use:   "list-member-groups",
	Short: "Returns a list of groups that a member belongs to within a tenant.",
	RunE:  impl.ListMemberGroups,
}

// listMemberPermissions -- Returns a set of permissions granted to the member within the tenant.

var listMemberPermissionsCmd = &cobra.Command{
	Use:   "list-member-permissions",
	Short: "Returns a set of permissions granted to the member within the tenant.",
	RunE:  impl.ListMemberPermissions,
}

// listMemberRoles -- Returns a set of roles that a given member holds within the tenant.

var listMemberRolesCmd = &cobra.Command{
	Use:   "list-member-roles",
	Short: "Returns a set of roles that a given member holds within the tenant.",
	RunE:  impl.ListMemberRoles,
}

// listMembers -- Returns a list of members in a given tenant.
var listMembersCmd = &cobra.Command{
	Use:   "list-members",
	Short: "Returns a list of members in a given tenant.",
	RunE:  impl.ListMembers,
}

// listPrincipals -- Returns the list of principals that the Identity service knows about.

var listPrincipalsCmd = &cobra.Command{
	Use:   "list-principals",
	Short: "Returns the list of principals that the Identity service knows about.",
	RunE:  impl.ListPrincipals,
}

// listRoleGroups -- Gets a list of groups for a role in a given tenant.
var listRoleGroupsCmd = &cobra.Command{
	Use:   "list-role-groups",
	Short: "Gets a list of groups for a role in a given tenant.",
	RunE:  impl.ListRoleGroups,
}

// listRolePermissions -- Gets the permissions for a role in a given tenant.
var listRolePermissionsCmd = &cobra.Command{
	Use:   "list-role-permissions",
	Short: "Gets the permissions for a role in a given tenant.",
	RunE:  impl.ListRolePermissions,
}

// listRoles -- Returns all roles for a given tenant.
var listRolesCmd = &cobra.Command{
	Use:   "list-roles",
	Short: "Returns all roles for a given tenant.",
	RunE:  impl.ListRoles,
}

// removeGroupMember -- Removes the member from a given group.
var removeGroupMemberCmd = &cobra.Command{
	Use:   "remove-group-member",
	Short: "Removes the member from a given group.",
	RunE:  impl.RemoveGroupMember,
}

// removeGroupRole -- Removes a role from a given group.
var removeGroupRoleCmd = &cobra.Command{
	Use:   "remove-group-role",
	Short: "Removes a role from a given group.",
	RunE:  impl.RemoveGroupRole,
}

// removeMember -- Removes a member from a given tenant
var removeMemberCmd = &cobra.Command{
	Use:   "remove-member",
	Short: "Removes a member from a given tenant",
	RunE:  impl.RemoveMember,
}

// removeRolePermission -- Removes a permission from the role.
var removeRolePermissionCmd = &cobra.Command{
	Use:   "remove-role-permission",
	Short: "Removes a permission from the role.",
	RunE:  impl.RemoveRolePermission,
}

// revokePrincipalAuthTokens -- Revoke all existing access tokens issued to a principal. Principals can reset their password by visiting https://login.splunk.com/en_us/page/lost_password

var revokePrincipalAuthTokensCmd = &cobra.Command{
	Use:   "revoke-principal-auth-tokens",
	Short: "Revoke all existing access tokens issued to a principal. Principals can reset their password by visiting https://login.splunk.com/en_us/page/lost_password",
	RunE:  impl.RevokePrincipalAuthTokens,
}

// updatePrincipalPublicKey -- Update principal public key
var updatePrincipalPublicKeyCmd = &cobra.Command{
	Use:   "update-principal-public-key",
	Short: "Update principal public key",
	RunE:  impl.UpdatePrincipalPublicKey,
}

// validateToken -- Validates the access token obtained from the authorization header and returns the principal name and tenant memberships.

var validateTokenCmd = &cobra.Command{
	Use:   "validate-token",
	Short: "Validates the access token obtained from the authorization header and returns the principal name and tenant memberships.",
	RunE:  impl.ValidateToken,
}

func init() {
	identityCmd.AddCommand(addGroupMemberCmd)

	var addGroupMemberGroup string
	addGroupMemberCmd.Flags().StringVar(&addGroupMemberGroup, "group", "", "This is a required parameter. The group name.")
	addGroupMemberCmd.MarkFlagRequired("group")

	var addGroupMemberName string
	addGroupMemberCmd.Flags().StringVar(&addGroupMemberName, "name", "", "This is a required parameter. ")
	addGroupMemberCmd.MarkFlagRequired("name")

	identityCmd.AddCommand(addGroupRoleCmd)

	var addGroupRoleGroup string
	addGroupRoleCmd.Flags().StringVar(&addGroupRoleGroup, "group", "", "This is a required parameter. The group name.")
	addGroupRoleCmd.MarkFlagRequired("group")

	var addGroupRoleName string
	addGroupRoleCmd.Flags().StringVar(&addGroupRoleName, "name", "", "This is a required parameter. ")
	addGroupRoleCmd.MarkFlagRequired("name")

	identityCmd.AddCommand(addMemberCmd)

	var addMemberName string
	addMemberCmd.Flags().StringVar(&addMemberName, "name", "", "This is a required parameter. ")
	addMemberCmd.MarkFlagRequired("name")

	identityCmd.AddCommand(addPrincipalPublicKeyCmd)

	var addPrincipalPublicKeyPrincipal string
	addPrincipalPublicKeyCmd.Flags().StringVar(&addPrincipalPublicKeyPrincipal, "principal", "", "This is a required parameter. The principal name.")
	addPrincipalPublicKeyCmd.MarkFlagRequired("principal")

	var addPrincipalPublicKeyAlg string
	addPrincipalPublicKeyCmd.Flags().StringVar(&addPrincipalPublicKeyAlg, "alg", "", " can accept values ES256, ES384, ES512")

	var addPrincipalPublicKeyCrv string
	addPrincipalPublicKeyCmd.Flags().StringVar(&addPrincipalPublicKeyCrv, "crv", "", "")

	var addPrincipalPublicKeyD string
	addPrincipalPublicKeyCmd.Flags().StringVar(&addPrincipalPublicKeyD, "d", "", "")

	var addPrincipalPublicKeyKid string
	addPrincipalPublicKeyCmd.Flags().StringVar(&addPrincipalPublicKeyKid, "kid", "", "")

	var addPrincipalPublicKeyKty string
	addPrincipalPublicKeyCmd.Flags().StringVar(&addPrincipalPublicKeyKty, "kty", "", " can accept values EC")

	var addPrincipalPublicKeyX string
	addPrincipalPublicKeyCmd.Flags().StringVar(&addPrincipalPublicKeyX, "x", "", "")

	var addPrincipalPublicKeyY string
	addPrincipalPublicKeyCmd.Flags().StringVar(&addPrincipalPublicKeyY, "y", "", "")

	identityCmd.AddCommand(addRolePermissionCmd)

	var addRolePermissionPermission string
	addRolePermissionCmd.Flags().StringVar(&addRolePermissionPermission, "permission", "", "This is a required parameter. ")
	addRolePermissionCmd.MarkFlagRequired("permission")

	var addRolePermissionRole string
	addRolePermissionCmd.Flags().StringVar(&addRolePermissionRole, "role", "", "This is a required parameter. The role name.")
	addRolePermissionCmd.MarkFlagRequired("role")

	identityCmd.AddCommand(createGroupCmd)

	var createGroupName string
	createGroupCmd.Flags().StringVar(&createGroupName, "name", "", "This is a required parameter. ")
	createGroupCmd.MarkFlagRequired("name")

	identityCmd.AddCommand(createPrincipalCmd)

	var createPrincipalEmail string
	createPrincipalCmd.Flags().StringVar(&createPrincipalEmail, "email", "", "This is a required parameter. ")
	createPrincipalCmd.MarkFlagRequired("email")

	var createPrincipalFirstName string
	createPrincipalCmd.Flags().StringVar(&createPrincipalFirstName, "first-name", "", "This is a required parameter. ")
	createPrincipalCmd.MarkFlagRequired("first-name")

	var createPrincipalKind string
	createPrincipalCmd.Flags().StringVar(&createPrincipalKind, "kind", "", "This is a required parameter.  can accept values user, service_account, service")
	createPrincipalCmd.MarkFlagRequired("kind")

	var createPrincipalLastName string
	createPrincipalCmd.Flags().StringVar(&createPrincipalLastName, "last-name", "", "This is a required parameter. ")
	createPrincipalCmd.MarkFlagRequired("last-name")

	var createPrincipalAlg string
	createPrincipalCmd.Flags().StringVar(&createPrincipalAlg, "alg", "", " can accept values ES256, ES384, ES512")

	var createPrincipalCredentials string
	createPrincipalCmd.Flags().StringVar(&createPrincipalCredentials, "credentials", "", "List of credentials")

	var createPrincipalCrv string
	createPrincipalCmd.Flags().StringVar(&createPrincipalCrv, "crv", "", "")

	var createPrincipalD string
	createPrincipalCmd.Flags().StringVar(&createPrincipalD, "d", "", "")

	var createPrincipalEnabled string
	createPrincipalCmd.Flags().StringVar(&createPrincipalEnabled, "enabled", "false", "")

	var createPrincipalInviteId string
	createPrincipalCmd.Flags().StringVar(&createPrincipalInviteId, "invite-id", "", "The invite ID.")

	var createPrincipalKid string
	createPrincipalCmd.Flags().StringVar(&createPrincipalKid, "kid", "", "")

	var createPrincipalKty string
	createPrincipalCmd.Flags().StringVar(&createPrincipalKty, "kty", "", " can accept values EC")

	var createPrincipalName string
	createPrincipalCmd.Flags().StringVar(&createPrincipalName, "name", "", "")

	var createPrincipalX string
	createPrincipalCmd.Flags().StringVar(&createPrincipalX, "x", "", "")

	var createPrincipalY string
	createPrincipalCmd.Flags().StringVar(&createPrincipalY, "y", "", "")

	identityCmd.AddCommand(createRoleCmd)

	var createRoleName string
	createRoleCmd.Flags().StringVar(&createRoleName, "name", "", "This is a required parameter. ")
	createRoleCmd.MarkFlagRequired("name")

	identityCmd.AddCommand(deleteGroupCmd)

	var deleteGroupGroup string
	deleteGroupCmd.Flags().StringVar(&deleteGroupGroup, "group", "", "This is a required parameter. The group name.")
	deleteGroupCmd.MarkFlagRequired("group")

	identityCmd.AddCommand(deletePrincipalPublicKeyCmd)

	var deletePrincipalPublicKeyKeyId string
	deletePrincipalPublicKeyCmd.Flags().StringVar(&deletePrincipalPublicKeyKeyId, "key-id", "", "This is a required parameter. Identifier of a public key.")
	deletePrincipalPublicKeyCmd.MarkFlagRequired("key-id")

	var deletePrincipalPublicKeyPrincipal string
	deletePrincipalPublicKeyCmd.Flags().StringVar(&deletePrincipalPublicKeyPrincipal, "principal", "", "This is a required parameter. The principal name.")
	deletePrincipalPublicKeyCmd.MarkFlagRequired("principal")

	identityCmd.AddCommand(deleteRoleCmd)

	var deleteRoleRole string
	deleteRoleCmd.Flags().StringVar(&deleteRoleRole, "role", "", "This is a required parameter. The role name.")
	deleteRoleCmd.MarkFlagRequired("role")

	identityCmd.AddCommand(getGroupCmd)

	var getGroupGroup string
	getGroupCmd.Flags().StringVar(&getGroupGroup, "group", "", "This is a required parameter. The group name.")
	getGroupCmd.MarkFlagRequired("group")

	identityCmd.AddCommand(getGroupMemberCmd)

	var getGroupMemberGroup string
	getGroupMemberCmd.Flags().StringVar(&getGroupMemberGroup, "group", "", "This is a required parameter. The group name.")
	getGroupMemberCmd.MarkFlagRequired("group")

	var getGroupMemberMember string
	getGroupMemberCmd.Flags().StringVar(&getGroupMemberMember, "member", "", "This is a required parameter. The member name.")
	getGroupMemberCmd.MarkFlagRequired("member")

	identityCmd.AddCommand(getGroupRoleCmd)

	var getGroupRoleGroup string
	getGroupRoleCmd.Flags().StringVar(&getGroupRoleGroup, "group", "", "This is a required parameter. The group name.")
	getGroupRoleCmd.MarkFlagRequired("group")

	var getGroupRoleRole string
	getGroupRoleCmd.Flags().StringVar(&getGroupRoleRole, "role", "", "This is a required parameter. The role name.")
	getGroupRoleCmd.MarkFlagRequired("role")

	identityCmd.AddCommand(getMemberCmd)

	var getMemberMember string
	getMemberCmd.Flags().StringVar(&getMemberMember, "member", "", "This is a required parameter. The member name.")
	getMemberCmd.MarkFlagRequired("member")

	identityCmd.AddCommand(getPrincipalCmd)

	var getPrincipalPrincipal string
	getPrincipalCmd.Flags().StringVar(&getPrincipalPrincipal, "principal", "", "This is a required parameter. The principal name.")
	getPrincipalCmd.MarkFlagRequired("principal")

	identityCmd.AddCommand(getPrincipalPublicKeyCmd)

	var getPrincipalPublicKeyKeyId string
	getPrincipalPublicKeyCmd.Flags().StringVar(&getPrincipalPublicKeyKeyId, "key-id", "", "This is a required parameter. Identifier of a public key.")
	getPrincipalPublicKeyCmd.MarkFlagRequired("key-id")

	var getPrincipalPublicKeyPrincipal string
	getPrincipalPublicKeyCmd.Flags().StringVar(&getPrincipalPublicKeyPrincipal, "principal", "", "This is a required parameter. The principal name.")
	getPrincipalPublicKeyCmd.MarkFlagRequired("principal")

	identityCmd.AddCommand(getPrincipalPublicKeysCmd)

	var getPrincipalPublicKeysPrincipal string
	getPrincipalPublicKeysCmd.Flags().StringVar(&getPrincipalPublicKeysPrincipal, "principal", "", "This is a required parameter. The principal name.")
	getPrincipalPublicKeysCmd.MarkFlagRequired("principal")

	identityCmd.AddCommand(getRoleCmd)

	var getRoleRole string
	getRoleCmd.Flags().StringVar(&getRoleRole, "role", "", "This is a required parameter. The role name.")
	getRoleCmd.MarkFlagRequired("role")

	identityCmd.AddCommand(getRolePermissionCmd)

	var getRolePermissionPermission string
	getRolePermissionCmd.Flags().StringVar(&getRolePermissionPermission, "permission", "", "This is a required parameter. The permission string.")
	getRolePermissionCmd.MarkFlagRequired("permission")

	var getRolePermissionRole string
	getRolePermissionCmd.Flags().StringVar(&getRolePermissionRole, "role", "", "This is a required parameter. The role name.")
	getRolePermissionCmd.MarkFlagRequired("role")

	identityCmd.AddCommand(listGroupMembersCmd)

	var listGroupMembersGroup string
	listGroupMembersCmd.Flags().StringVar(&listGroupMembersGroup, "group", "", "This is a required parameter. The group name.")
	listGroupMembersCmd.MarkFlagRequired("group")

	var listGroupMembersOrderby string
	listGroupMembersCmd.Flags().StringVar(&listGroupMembersOrderby, "orderby", "", "The sorting order for returning list.")

	var listGroupMembersPageSize int32
	listGroupMembersCmd.Flags().Int32Var(&listGroupMembersPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listGroupMembersPageToken string
	listGroupMembersCmd.Flags().StringVar(&listGroupMembersPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(listGroupRolesCmd)

	var listGroupRolesGroup string
	listGroupRolesCmd.Flags().StringVar(&listGroupRolesGroup, "group", "", "This is a required parameter. The group name.")
	listGroupRolesCmd.MarkFlagRequired("group")

	var listGroupRolesOrderby string
	listGroupRolesCmd.Flags().StringVar(&listGroupRolesOrderby, "orderby", "", "The sorting order for returning list.")

	var listGroupRolesPageSize int32
	listGroupRolesCmd.Flags().Int32Var(&listGroupRolesPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listGroupRolesPageToken string
	listGroupRolesCmd.Flags().StringVar(&listGroupRolesPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(listGroupsCmd)

	var listGroupsAccess string
	listGroupsCmd.Flags().StringVar(&listGroupsAccess, "access", "", "List only the groups with specified access permission.")

	var listGroupsOrderby string
	listGroupsCmd.Flags().StringVar(&listGroupsOrderby, "orderby", "", "The sorting order for returning list.")

	var listGroupsPageSize int32
	listGroupsCmd.Flags().Int32Var(&listGroupsPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listGroupsPageToken string
	listGroupsCmd.Flags().StringVar(&listGroupsPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(listMemberGroupsCmd)

	var listMemberGroupsMember string
	listMemberGroupsCmd.Flags().StringVar(&listMemberGroupsMember, "member", "", "This is a required parameter. The member name.")
	listMemberGroupsCmd.MarkFlagRequired("member")

	var listMemberGroupsOrderby string
	listMemberGroupsCmd.Flags().StringVar(&listMemberGroupsOrderby, "orderby", "", "The sorting order for returning list.")

	var listMemberGroupsPageSize int32
	listMemberGroupsCmd.Flags().Int32Var(&listMemberGroupsPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listMemberGroupsPageToken string
	listMemberGroupsCmd.Flags().StringVar(&listMemberGroupsPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(listMemberPermissionsCmd)

	var listMemberPermissionsMember string
	listMemberPermissionsCmd.Flags().StringVar(&listMemberPermissionsMember, "member", "", "This is a required parameter. The member name.")
	listMemberPermissionsCmd.MarkFlagRequired("member")

	var listMemberPermissionsOrderby string
	listMemberPermissionsCmd.Flags().StringVar(&listMemberPermissionsOrderby, "orderby", "", "The sorting order for returning list.")

	var listMemberPermissionsPageSize int32
	listMemberPermissionsCmd.Flags().Int32Var(&listMemberPermissionsPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listMemberPermissionsPageToken string
	listMemberPermissionsCmd.Flags().StringVar(&listMemberPermissionsPageToken, "page-token", "", "The cursor to then next page.")

	var listMemberPermissionsScopeFilter string
	listMemberPermissionsCmd.Flags().StringVar(&listMemberPermissionsScopeFilter, "scope-filter", "", "List only the permissions matching the scope filter.")

	identityCmd.AddCommand(listMemberRolesCmd)

	var listMemberRolesMember string
	listMemberRolesCmd.Flags().StringVar(&listMemberRolesMember, "member", "", "This is a required parameter. The member name.")
	listMemberRolesCmd.MarkFlagRequired("member")

	var listMemberRolesOrderby string
	listMemberRolesCmd.Flags().StringVar(&listMemberRolesOrderby, "orderby", "", "The sorting order for returning list.")

	var listMemberRolesPageSize int32
	listMemberRolesCmd.Flags().Int32Var(&listMemberRolesPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listMemberRolesPageToken string
	listMemberRolesCmd.Flags().StringVar(&listMemberRolesPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(listMembersCmd)

	var listMembersKind string
	listMembersCmd.Flags().StringVar(&listMembersKind, "kind", "", "Kind of member, one of: [user, service_account, service]")

	var listMembersOrderby string
	listMembersCmd.Flags().StringVar(&listMembersOrderby, "orderby", "", "The sorting order for returning list.")

	var listMembersPageSize int32
	listMembersCmd.Flags().Int32Var(&listMembersPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listMembersPageToken string
	listMembersCmd.Flags().StringVar(&listMembersPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(listPrincipalsCmd)

	var listPrincipalsOrderby string
	listPrincipalsCmd.Flags().StringVar(&listPrincipalsOrderby, "orderby", "", "The sorting order for returning list.")

	var listPrincipalsPageSize int32
	listPrincipalsCmd.Flags().Int32Var(&listPrincipalsPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listPrincipalsPageToken string
	listPrincipalsCmd.Flags().StringVar(&listPrincipalsPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(listRoleGroupsCmd)

	var listRoleGroupsRole string
	listRoleGroupsCmd.Flags().StringVar(&listRoleGroupsRole, "role", "", "This is a required parameter. The role name.")
	listRoleGroupsCmd.MarkFlagRequired("role")

	var listRoleGroupsOrderby string
	listRoleGroupsCmd.Flags().StringVar(&listRoleGroupsOrderby, "orderby", "", "The sorting order for returning list.")

	var listRoleGroupsPageSize int32
	listRoleGroupsCmd.Flags().Int32Var(&listRoleGroupsPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listRoleGroupsPageToken string
	listRoleGroupsCmd.Flags().StringVar(&listRoleGroupsPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(listRolePermissionsCmd)

	var listRolePermissionsRole string
	listRolePermissionsCmd.Flags().StringVar(&listRolePermissionsRole, "role", "", "This is a required parameter. The role name.")
	listRolePermissionsCmd.MarkFlagRequired("role")

	var listRolePermissionsOrderby string
	listRolePermissionsCmd.Flags().StringVar(&listRolePermissionsOrderby, "orderby", "", "The sorting order for returning list.")

	var listRolePermissionsPageSize int32
	listRolePermissionsCmd.Flags().Int32Var(&listRolePermissionsPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listRolePermissionsPageToken string
	listRolePermissionsCmd.Flags().StringVar(&listRolePermissionsPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(listRolesCmd)

	var listRolesOrderby string
	listRolesCmd.Flags().StringVar(&listRolesOrderby, "orderby", "", "The sorting order for returning list.")

	var listRolesPageSize int32
	listRolesCmd.Flags().Int32Var(&listRolesPageSize, "page-size", 0, "The maximize return items count of a list.")

	var listRolesPageToken string
	listRolesCmd.Flags().StringVar(&listRolesPageToken, "page-token", "", "The cursor to then next page.")

	identityCmd.AddCommand(removeGroupMemberCmd)

	var removeGroupMemberGroup string
	removeGroupMemberCmd.Flags().StringVar(&removeGroupMemberGroup, "group", "", "This is a required parameter. The group name.")
	removeGroupMemberCmd.MarkFlagRequired("group")

	var removeGroupMemberMember string
	removeGroupMemberCmd.Flags().StringVar(&removeGroupMemberMember, "member", "", "This is a required parameter. The member name.")
	removeGroupMemberCmd.MarkFlagRequired("member")

	identityCmd.AddCommand(removeGroupRoleCmd)

	var removeGroupRoleGroup string
	removeGroupRoleCmd.Flags().StringVar(&removeGroupRoleGroup, "group", "", "This is a required parameter. The group name.")
	removeGroupRoleCmd.MarkFlagRequired("group")

	var removeGroupRoleRole string
	removeGroupRoleCmd.Flags().StringVar(&removeGroupRoleRole, "role", "", "This is a required parameter. The role name.")
	removeGroupRoleCmd.MarkFlagRequired("role")

	identityCmd.AddCommand(removeMemberCmd)

	var removeMemberMember string
	removeMemberCmd.Flags().StringVar(&removeMemberMember, "member", "", "This is a required parameter. The member name.")
	removeMemberCmd.MarkFlagRequired("member")

	identityCmd.AddCommand(removeRolePermissionCmd)

	var removeRolePermissionPermission string
	removeRolePermissionCmd.Flags().StringVar(&removeRolePermissionPermission, "permission", "", "This is a required parameter. The permission string.")
	removeRolePermissionCmd.MarkFlagRequired("permission")

	var removeRolePermissionRole string
	removeRolePermissionCmd.Flags().StringVar(&removeRolePermissionRole, "role", "", "This is a required parameter. The role name.")
	removeRolePermissionCmd.MarkFlagRequired("role")

	identityCmd.AddCommand(revokePrincipalAuthTokensCmd)

	var revokePrincipalAuthTokensPrincipal string
	revokePrincipalAuthTokensCmd.Flags().StringVar(&revokePrincipalAuthTokensPrincipal, "principal", "", "This is a required parameter. The principal name.")
	revokePrincipalAuthTokensCmd.MarkFlagRequired("principal")

	identityCmd.AddCommand(updatePrincipalPublicKeyCmd)

	var updatePrincipalPublicKeyKeyId string
	updatePrincipalPublicKeyCmd.Flags().StringVar(&updatePrincipalPublicKeyKeyId, "key-id", "", "This is a required parameter. Identifier of a public key.")
	updatePrincipalPublicKeyCmd.MarkFlagRequired("key-id")

	var updatePrincipalPublicKeyPrincipal string
	updatePrincipalPublicKeyCmd.Flags().StringVar(&updatePrincipalPublicKeyPrincipal, "principal", "", "This is a required parameter. The principal name.")
	updatePrincipalPublicKeyCmd.MarkFlagRequired("principal")

	var updatePrincipalPublicKeyStatus string
	updatePrincipalPublicKeyCmd.Flags().StringVar(&updatePrincipalPublicKeyStatus, "status", "", "This is a required parameter.  can accept values active, inactive")
	updatePrincipalPublicKeyCmd.MarkFlagRequired("status")

	identityCmd.AddCommand(validateTokenCmd)

	var validateTokenInclude string
	validateTokenCmd.Flags().StringVar(&validateTokenInclude, "include", "", "Include additional information to return when validating tenant membership. Valid parameters [tenant, principal]")

}
