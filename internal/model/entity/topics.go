// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Topics is the golang structure for table topics.
type Topics struct {
	TopicId     uint        `json:"topic_id"    description:"话题ID"`
	Name        string      `json:"name"        description:"话题名称"`
	Description string      `json:"description" description:"话题描述"`
	Icon        string      `json:"icon"        description:"话题图标"`
	PostCount   uint        `json:"post_count"  description:"帖子数"`
	Status      int         `json:"status"      description:"状态：0-已删除，1-正常"`
	IsHot       int         `json:"is_hot"      description:"是否热门：0-否，1-是"`
	CreateTime  *gtime.Time `json:"create_time" description:"创建时间"`
	UpdateTime  *gtime.Time `json:"update_time" description:"更新时间"`
}
