// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogComments is the golang structure of table blog_comments for DAO operations like Where/Data.
type BlogComments struct {
	g.Meta     `orm:"table:blog_comments, do:true"`
	CommentId  interface{} // 评论ID，主键，自增
	BlogId     interface{} // 博客ID，关联blogs表
	UserId     interface{} // 评论用户ID，关联users表
	Content    interface{} // 评论内容
	ParentId   interface{} // 父评论ID，用于回复功能，如为NULL则为顶级评论
	LikeCount  interface{} // 点赞数
	Status     interface{} // 状态：0-已删除，1-正常
	CreateTime *gtime.Time // 记录创建时间
	UpdateTime *gtime.Time // 记录最后更新时间
}
