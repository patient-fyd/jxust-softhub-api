// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Comments is the golang structure of table comments for DAO operations like Where/Data.
type Comments struct {
	g.Meta      `orm:"table:comments, do:true"`
	CommentId   interface{} // 评论ID，主键，自增
	ContentType interface{} // 评论关联内容类型，如 "news" 或 "activity"
	ContentId   interface{} // 关联内容ID，指向具体的新闻或活动
	UserId      interface{} // 评论用户ID，关联 users 表
	Content     interface{} // 评论内容
	CreateTime  *gtime.Time // 记录创建时间
	LikeCount   interface{} // 点赞数
}
