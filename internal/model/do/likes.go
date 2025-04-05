// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Likes is the golang structure of table likes for DAO operations like Where/Data.
type Likes struct {
	g.Meta     `orm:"table:likes, do:true"`
	LikeId     interface{} // 点赞ID
	UserId     interface{} // 用户ID
	TargetId   interface{} // 目标ID
	TargetType interface{} // 目标类型：1-帖子，2-评论
	CreateTime *gtime.Time // 创建时间
}
