// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemConfigsDao is the data access object for table system_configs.
type SystemConfigsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns SystemConfigsColumns // columns contains all the column names of Table for convenient usage.
}

// SystemConfigsColumns defines and stores column names for table system_configs.
type SystemConfigsColumns struct {
	ConfigKey   string // 配置键名
	ConfigValue string // 配置值
	Description string // 配置说明
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
}

// systemConfigsColumns holds the columns for table system_configs.
var systemConfigsColumns = SystemConfigsColumns{
	ConfigKey:   "configKey",
	ConfigValue: "configValue",
	Description: "description",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
}

// NewSystemConfigsDao creates and returns a new DAO object for table data access.
func NewSystemConfigsDao() *SystemConfigsDao {
	return &SystemConfigsDao{
		group:   "default",
		table:   "system_configs",
		columns: systemConfigsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemConfigsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemConfigsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemConfigsDao) Columns() SystemConfigsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemConfigsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemConfigsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemConfigsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
