// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Activities is the golang structure for table activities.
type Activities struct {
	ActivityId      uint        `json:"activity_id"      description:"活动ID，主键，自增"`
	Title           string      `json:"title"            description:"活动标题"`
	Description     string      `json:"description"      description:"活动详细描述"`
	StartTime       *gtime.Time `json:"start_time"       description:"活动开始时间"`
	EndTime         *gtime.Time `json:"end_time"         description:"活动结束时间"`
	Location        string      `json:"location"         description:"活动举办地点"`
	MaxParticipants int         `json:"max_participants" description:"最大参与人数限制"`
	Status          int         `json:"status"           description:"活动状态：0-未开始, 1-进行中, 2-已结束"`
	CreateTime      *gtime.Time `json:"create_time"      description:"记录创建时间"`
	UpdateTime      *gtime.Time `json:"update_time"      description:"记录最后更新时间"`
}
