// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Activities is the golang structure of table activities for DAO operations like Where/Data.
type Activities struct {
	g.Meta          `orm:"table:activities, do:true"`
	ActivityId      interface{} // 活动ID，主键，自增
	Title           interface{} // 活动标题
	Description     interface{} // 活动详细描述
	StartTime       *gtime.Time // 活动开始时间
	EndTime         *gtime.Time // 活动结束时间
	Location        interface{} // 活动举办地点
	MaxParticipants interface{} // 最大参与人数限制
	Status          interface{} // 活动状态：0-未开始, 1-进行中, 2-已结束
	CreateTime      *gtime.Time // 记录创建时间
	UpdateTime      *gtime.Time // 记录最后更新时间
}
