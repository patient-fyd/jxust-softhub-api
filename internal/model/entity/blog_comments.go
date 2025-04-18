// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogComments is the golang structure for table blog_comments.
type BlogComments struct {
	CommentId  uint        `json:"commentId"  description:"评论ID，主键，自增"`
	BlogId     uint        `json:"blogId"     description:"博客ID，关联blogs表"`
	UserId     uint        `json:"userId"     description:"评论用户ID，关联users表"`
	Content    string      `json:"content"     description:"评论内容"`
	ParentId   uint        `json:"parentId"   description:"父评论ID，用于回复功能，如为NULL则为顶级评论"`
	LikeCount  uint        `json:"likeCount"  description:"点赞数"`
	Status     int         `json:"status"      description:"状态：0-已删除，1-正常"`
	CreateTime *gtime.Time `json:"createTime" description:"记录创建时间"`
	UpdateTime *gtime.Time `json:"updateTime" description:"记录最后更新时间"`
}
