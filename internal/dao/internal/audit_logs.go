// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuditLogsDao is the data access object for table audit_logs.
type AuditLogsDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns AuditLogsColumns // columns contains all the column names of Table for convenient usage.
}

// AuditLogsColumns defines and stores column names for table audit_logs.
type AuditLogsColumns struct {
	LogId       string // 日志ID，主键，自增
	UserId      string // 操作用户ID，关联 users 表，如为空表示系统自动操作
	Action      string // 操作名称或类型，如 "login", "update_news" 等
	Description string // 操作详细描述，记录关键信息
	CreateTime  string // 记录操作时间
}

// auditLogsColumns holds the columns for table audit_logs.
var auditLogsColumns = AuditLogsColumns{
	LogId:       "logId",
	UserId:      "userId",
	Action:      "action",
	Description: "description",
	CreateTime:  "createTime",
}

// NewAuditLogsDao creates and returns a new DAO object for table data access.
func NewAuditLogsDao() *AuditLogsDao {
	return &AuditLogsDao{
		group:   "default",
		table:   "audit_logs",
		columns: auditLogsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AuditLogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AuditLogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AuditLogsDao) Columns() AuditLogsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AuditLogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AuditLogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AuditLogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
