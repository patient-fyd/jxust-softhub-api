// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Members is the golang structure for table members.
type Members struct {
	MemberId     uint        `json:"member_id"    description:"成员ID，主键，自增"`
	UserId       uint        `json:"user_id"      description:"关联的用户ID"`
	Grade        string      `json:"grade"        description:"年级，如2020级"`
	JoinYear     string      `json:"join_year"    description:"入会年份"`
	Department   string      `json:"department"   description:"所属部门，如技术部、宣传部等"`
	Position     string      `json:"position"     description:"职位，如部长、副部长、干事等"`
	Skills       string      `json:"skills"       description:"技能描述"`
	Introduction string      `json:"introduction" description:"个人简介"`
	IsCore       int         `json:"is_core"      description:"是否为核心成员：0-否，1-是"`
	Status       int         `json:"status"       description:"成员状态：0-待审核，1-正式成员，2-已退出"`
	CreateTime   *gtime.Time `json:"create_time"  description:"记录创建时间"`
	UpdateTime   *gtime.Time `json:"update_time"  description:"记录最后更新时间"`
}
