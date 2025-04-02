// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Permissions is the golang structure for table permissions.
type Permissions struct {
	PermissionId  uint        `json:"permission_id"  description:"权限ID，主键，自增"`
	PermissionKey string      `json:"permission_key" description:"权限标识，用于后台校验，如 \"news_edit\""`
	Description   string      `json:"description"    description:"权限描述信息"`
	CreateTime    *gtime.Time `json:"create_time"    description:"记录创建时间"`
	UpdateTime    *gtime.Time `json:"update_time"    description:"记录最后更新时间"`
}
