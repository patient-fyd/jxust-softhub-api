// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogLikes is the golang structure for table blog_likes.
type BlogLikes struct {
	LikeId     uint        `json:"like_id"     description:"点赞ID，主键，自增"`
	BlogId     uint        `json:"blog_id"     description:"博客ID，关联blogs表"`
	UserId     uint        `json:"user_id"     description:"用户ID，关联users表"`
	CreateTime *gtime.Time `json:"create_time" description:"点赞时间"`
}
