// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Follows is the golang structure of table follows for DAO operations like Where/Data.
type Follows struct {
	g.Meta     `orm:"table:follows, do:true"`
	FollowId   interface{} // 关注ID
	UserId     interface{} // 用户ID
	FollowedId interface{} // 被关注对象ID
	FollowType interface{} // 关注类型：1-用户，2-圈子，3-话题
	CreateTime *gtime.Time // 创建时间
}
