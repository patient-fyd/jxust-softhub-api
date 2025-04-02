// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure of table permissions for DAO operations like Where/Data.
type Permissions struct {
	g.Meta        `orm:"table:permissions, do:true"`
	PermissionId  interface{} // 权限ID，主键，自增
	PermissionKey interface{} // 权限标识，用于后台校验，如 "news_edit"
	Description   interface{} // 权限描述信息
	CreateTime    *gtime.Time // 记录创建时间
	UpdateTime    *gtime.Time // 记录最后更新时间
}
