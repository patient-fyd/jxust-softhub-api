// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Follows is the golang structure for table follows.
type Follows struct {
	FollowId   uint        `json:"follow_id"   description:"关注ID"`
	UserId     uint        `json:"user_id"     description:"用户ID"`
	FollowedId uint        `json:"followed_id" description:"被关注对象ID"`
	FollowType int         `json:"follow_type" description:"关注类型：1-用户，2-圈子，3-话题"`
	CreateTime *gtime.Time `json:"create_time" description:"创建时间"`
}
