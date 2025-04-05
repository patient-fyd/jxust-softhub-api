// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Posts is the golang structure for table posts.
type Posts struct {
	PostId       uint        `json:"post_id"       description:"帖子ID"`
	UserId       uint        `json:"user_id"       description:"发帖用户ID"`
	Content      string      `json:"content"       description:"帖子内容"`
	CircleId     uint        `json:"circle_id"     description:"所属圈子ID"`
	TopicId      uint        `json:"topic_id"      description:"所属话题ID"`
	ViewCount    uint        `json:"view_count"    description:"浏览量"`
	LikeCount    uint        `json:"like_count"    description:"点赞数"`
	CommentCount uint        `json:"comment_count" description:"评论数"`
	ShareCount   uint        `json:"share_count"   description:"分享数"`
	IsTop        int         `json:"is_top"        description:"是否置顶：0-否，1-是"`
	Status       int         `json:"status"        description:"状态：0-草稿，1-已发布，2-已删除"`
	CreateTime   *gtime.Time `json:"create_time"   description:"创建时间"`
	UpdateTime   *gtime.Time `json:"update_time"   description:"更新时间"`
}
