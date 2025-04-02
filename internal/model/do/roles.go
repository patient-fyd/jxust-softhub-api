// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Roles is the golang structure of table roles for DAO operations like Where/Data.
type Roles struct {
	g.Meta      `orm:"table:roles, do:true"`
	RoleId      interface{} // 角色ID，主键，自增
	RoleName    interface{} // 角色名称，如超级管理员、内容管理员等
	Description interface{} // 角色描述信息
	CreateTime  *gtime.Time // 记录创建时间
	UpdateTime  *gtime.Time // 记录最后更新时间
}
