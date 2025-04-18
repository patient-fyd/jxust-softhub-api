// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogLikes is the golang structure for table blog_likes.
type BlogLikes struct {
	LikeId     uint        `json:"likeId"     description:"点赞ID，主键，自增"`
	BlogId     uint        `json:"blogId"     description:"博客ID，关联blogs表"`
	UserId     uint        `json:"userId"     description:"用户ID，关联users表"`
	CreateTime *gtime.Time `json:"createTime" description:"点赞时间"`
}
