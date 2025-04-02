// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuditLogs is the golang structure for table audit_logs.
type AuditLogs struct {
	LogId       uint        `json:"log_id"      description:"日志ID，主键，自增"`
	UserId      uint        `json:"user_id"     description:"操作用户ID，关联 users 表，如为空表示系统自动操作"`
	Action      string      `json:"action"      description:"操作名称或类型，如 \"login\", \"update_news\" 等"`
	Description string      `json:"description" description:"操作详细描述，记录关键信息"`
	CreateTime  *gtime.Time `json:"create_time" description:"记录操作时间"`
}
