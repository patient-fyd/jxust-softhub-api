// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// RolePermissions is the golang structure for table role_permissions.
type RolePermissions struct {
	RoleId       uint `json:"role_id"       description:"角色ID，关联 roles 表"`
	PermissionId uint `json:"permission_id" description:"权限ID，关联 permissions 表"`
}
