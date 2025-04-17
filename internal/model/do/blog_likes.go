// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BlogLikes is the golang structure of table blog_likes for DAO operations like Where/Data.
type BlogLikes struct {
	g.Meta     `orm:"table:blog_likes, do:true"`
	LikeId     interface{} // 点赞ID，主键，自增
	BlogId     interface{} // 博客ID，关联blogs表
	UserId     interface{} // 用户ID，关联users表
	CreateTime *gtime.Time // 点赞时间
}
