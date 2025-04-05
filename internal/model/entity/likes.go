// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Likes is the golang structure for table likes.
type Likes struct {
	LikeId     uint        `json:"like_id"     description:"点赞ID"`
	UserId     uint        `json:"user_id"     description:"用户ID"`
	TargetId   uint        `json:"target_id"   description:"目标ID"`
	TargetType int         `json:"target_type" description:"目标类型：1-帖子，2-评论"`
	CreateTime *gtime.Time `json:"create_time" description:"创建时间"`
}
