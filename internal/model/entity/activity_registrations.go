// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ActivityRegistrations is the golang structure for table activity_registrations.
type ActivityRegistrations struct {
	RegistrationId uint        `json:"registration_id" description:"报名记录ID，主键，自增"`
	ActivityId     uint        `json:"activity_id"     description:"活动ID，关联 activities 表"`
	UserId         uint        `json:"user_id"         description:"报名用户ID，关联 users 表，如为空表示未登录报名"`
	Name           string      `json:"name"            description:"报名者姓名"`
	StudentId      string      `json:"student_id"      description:"报名者学号"`
	Contact        string      `json:"contact"         description:"报名者联系方式，如电话或邮箱"`
	Status         int         `json:"status"          description:"报名状态：0-待审核, 1-通过, 2-拒绝"`
	CreateTime     *gtime.Time `json:"create_time"     description:"记录创建时间"`
	UpdateTime     *gtime.Time `json:"update_time"     description:"记录最后更新时间"`
}
