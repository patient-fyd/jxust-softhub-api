// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Circles is the golang structure of table circles for DAO operations like Where/Data.
type Circles struct {
	g.Meta      `orm:"table:circles, do:true"`
	CircleId    interface{} // 圈子ID
	Name        interface{} // 圈子名称
	Description interface{} // 圈子描述
	Icon        interface{} // 圈子图标
	UserId      interface{} // 创建者ID
	PostCount   interface{} // 帖子数
	MemberCount interface{} // 成员数
	Status      interface{} // 状态：0-已删除，1-正常
	IsOfficial  interface{} // 是否官方圈子：0-否，1-是
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
