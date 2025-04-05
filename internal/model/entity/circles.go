// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Circles is the golang structure for table circles.
type Circles struct {
	CircleId    uint        `json:"circle_id"    description:"圈子ID"`
	Name        string      `json:"name"         description:"圈子名称"`
	Description string      `json:"description"  description:"圈子描述"`
	Icon        string      `json:"icon"         description:"圈子图标"`
	UserId      uint        `json:"user_id"      description:"创建者ID"`
	PostCount   uint        `json:"post_count"   description:"帖子数"`
	MemberCount uint        `json:"member_count" description:"成员数"`
	Status      int         `json:"status"       description:"状态：0-已删除，1-正常"`
	IsOfficial  int         `json:"is_official"  description:"是否官方圈子：0-否，1-是"`
	CreateTime  *gtime.Time `json:"create_time"  description:"创建时间"`
	UpdateTime  *gtime.Time `json:"update_time"  description:"更新时间"`
}
