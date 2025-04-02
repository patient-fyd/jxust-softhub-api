// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// RolePermissions is the golang structure of table role_permissions for DAO operations like Where/Data.
type RolePermissions struct {
	g.Meta       `orm:"table:role_permissions, do:true"`
	RoleId       interface{} // 角色ID，关联 roles 表
	PermissionId interface{} // 权限ID，关联 permissions 表
}
