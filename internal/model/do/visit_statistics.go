// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VisitStatistics is the golang structure of table visit_statistics for DAO operations like Where/Data.
type VisitStatistics struct {
	g.Meta        `orm:"table:visit_statistics, do:true"`
	StatId        interface{} // 统计ID，主键，自增
	VisitDate     *gtime.Time // 访问日期
	PageView      interface{} // 页面浏览量
	UniqueVisitor interface{} // 独立访客数
	NewVisitor    interface{} // 新访客数
	AvgStayTime   interface{} // 平均停留时间（秒）
	BounceRate    interface{} // 跳出率（百分比）
	CreateTime    *gtime.Time // 记录创建时间
	UpdateTime    *gtime.Time // 记录最后更新时间
}
