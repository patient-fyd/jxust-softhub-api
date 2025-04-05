// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Comments is the golang structure for table comments.
type Comments struct {
	CommentId   uint        `json:"comment_id"   description:"评论ID，主键，自增"`
	ContentType string      `json:"content_type" description:"评论关联内容类型，如 \"news\" 或 \"activity\""`
	ContentId   uint        `json:"content_id"   description:"关联内容ID，指向具体的新闻或活动"`
	UserId      uint        `json:"user_id"      description:"评论用户ID，关联 users 表"`
	Content     string      `json:"content"      description:"评论内容"`
	CreateTime  *gtime.Time `json:"create_time"  description:"记录创建时间"`
	LikeCount   uint        `json:"like_count"   description:"点赞数"`
}
