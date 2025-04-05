// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Posts is the golang structure of table posts for DAO operations like Where/Data.
type Posts struct {
	g.Meta       `orm:"table:posts, do:true"`
	PostId       interface{} // 帖子ID
	UserId       interface{} // 发帖用户ID
	Content      interface{} // 帖子内容
	CircleId     interface{} // 所属圈子ID
	TopicId      interface{} // 所属话题ID
	ViewCount    interface{} // 浏览量
	LikeCount    interface{} // 点赞数
	CommentCount interface{} // 评论数
	ShareCount   interface{} // 分享数
	IsTop        interface{} // 是否置顶：0-否，1-是
	Status       interface{} // 状态：0-草稿，1-已发布，2-已删除
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
}
