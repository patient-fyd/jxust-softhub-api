// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogComments is the golang structure for table blog_comments.
type BlogComments struct {
	CommentId  uint        `json:"comment_id"  description:"评论ID，主键，自增"`
	BlogId     uint        `json:"blog_id"     description:"博客ID，关联blogs表"`
	UserId     uint        `json:"user_id"     description:"评论用户ID，关联users表"`
	Content    string      `json:"content"     description:"评论内容"`
	ParentId   uint        `json:"parent_id"   description:"父评论ID，用于回复功能，如为NULL则为顶级评论"`
	LikeCount  uint        `json:"like_count"  description:"点赞数"`
	Status     int         `json:"status"      description:"状态：0-已删除，1-正常"`
	CreateTime *gtime.Time `json:"create_time" description:"记录创建时间"`
	UpdateTime *gtime.Time `json:"update_time" description:"记录最后更新时间"`
}
