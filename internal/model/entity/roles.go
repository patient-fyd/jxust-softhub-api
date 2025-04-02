// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure for table roles.
type Roles struct {
	RoleId      uint        `json:"role_id"     description:"角色ID，主键，自增"`
	RoleName    string      `json:"role_name"   description:"角色名称，如超级管理员、内容管理员等"`
	Description string      `json:"description" description:"角色描述信息"`
	CreateTime  *gtime.Time `json:"create_time" description:"记录创建时间"`
	UpdateTime  *gtime.Time `json:"update_time" description:"记录最后更新时间"`
}
