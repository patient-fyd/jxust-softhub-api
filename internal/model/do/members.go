// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Members is the golang structure of table members for DAO operations like Where/Data.
type Members struct {
	g.Meta       `orm:"table:members, do:true"`
	MemberId     interface{} // 成员ID，主键，自增
	UserId       interface{} // 关联的用户ID
	Grade        interface{} // 年级，如2020级
	JoinYear     interface{} // 入会年份
	Department   interface{} // 所属部门，如技术部、宣传部等
	Position     interface{} // 职位，如部长、副部长、干事等
	Skills       interface{} // 技能描述
	Introduction interface{} // 个人简介
	IsCore       interface{} // 是否为核心成员：0-否，1-是
	Status       interface{} // 成员状态：0-待审核，1-正式成员，2-已退出
	CreateTime   *gtime.Time // 记录创建时间
	UpdateTime   *gtime.Time // 记录最后更新时间
}
