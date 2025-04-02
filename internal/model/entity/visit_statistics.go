// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitStatistics is the golang structure for table visit_statistics.
type VisitStatistics struct {
	StatId        uint        `json:"stat_id"        description:"统计ID，主键，自增"`
	VisitDate     *gtime.Time `json:"visit_date"     description:"访问日期"`
	PageView      uint        `json:"page_view"      description:"页面浏览量"`
	UniqueVisitor uint        `json:"unique_visitor" description:"独立访客数"`
	NewVisitor    uint        `json:"new_visitor"    description:"新访客数"`
	AvgStayTime   uint        `json:"avg_stay_time"  description:"平均停留时间（秒）"`
	BounceRate    float64     `json:"bounce_rate"    description:"跳出率（百分比）"`
	CreateTime    *gtime.Time `json:"create_time"    description:"记录创建时间"`
	UpdateTime    *gtime.Time `json:"update_time"    description:"记录最后更新时间"`
}
