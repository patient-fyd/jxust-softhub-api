// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Topics is the golang structure of table topics for DAO operations like Where/Data.
type Topics struct {
	g.Meta      `orm:"table:topics, do:true"`
	TopicId     interface{} // 话题ID
	Name        interface{} // 话题名称
	Description interface{} // 话题描述
	Icon        interface{} // 话题图标
	PostCount   interface{} // 帖子数
	Status      interface{} // 状态：0-已删除，1-正常
	IsHot       interface{} // 是否热门：0-否，1-是
	CreateTime  *gtime.Time // 创建时间
	UpdateTime  *gtime.Time // 更新时间
}
