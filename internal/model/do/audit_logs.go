// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuditLogs is the golang structure of table audit_logs for DAO operations like Where/Data.
type AuditLogs struct {
	g.Meta      `orm:"table:audit_logs, do:true"`
	LogId       interface{} // 日志ID，主键，自增
	UserId      interface{} // 操作用户ID，关联 users 表，如为空表示系统自动操作
	Action      interface{} // 操作名称或类型，如 "login", "update_news" 等
	Description interface{} // 操作详细描述，记录关键信息
	CreateTime  *gtime.Time // 记录操作时间
}
