// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityRegistrations is the golang structure of table activity_registrations for DAO operations like Where/Data.
type ActivityRegistrations struct {
	g.Meta         `orm:"table:activity_registrations, do:true"`
	RegistrationId interface{} // 报名记录ID，主键，自增
	ActivityId     interface{} // 活动ID，关联 activities 表
	UserId         interface{} // 报名用户ID，关联 users 表，如为空表示未登录报名
	Name           interface{} // 报名者姓名
	StudentId      interface{} // 报名者学号
	Contact        interface{} // 报名者联系方式，如电话或邮箱
	Status         interface{} // 报名状态：0-待审核, 1-通过, 2-拒绝
	CreateTime     *gtime.Time // 记录创建时间
	UpdateTime     *gtime.Time // 记录最后更新时间
}
