// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VisitStatisticsDao is the data access object for table visit_statistics.
type VisitStatisticsDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns VisitStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// VisitStatisticsColumns defines and stores column names for table visit_statistics.
type VisitStatisticsColumns struct {
	StatId        string // 统计ID，主键，自增
	VisitDate     string // 访问日期
	PageView      string // 页面浏览量
	UniqueVisitor string // 独立访客数
	NewVisitor    string // 新访客数
	AvgStayTime   string // 平均停留时间（秒）
	BounceRate    string // 跳出率（百分比）
	CreateTime    string // 记录创建时间
	UpdateTime    string // 记录最后更新时间
}

// visitStatisticsColumns holds the columns for table visit_statistics.
var visitStatisticsColumns = VisitStatisticsColumns{
	StatId:        "statId",
	VisitDate:     "visitDate",
	PageView:      "pageView",
	UniqueVisitor: "uniqueVisitor",
	NewVisitor:    "newVisitor",
	AvgStayTime:   "avgStayTime",
	BounceRate:    "bounceRate",
	CreateTime:    "createTime",
	UpdateTime:    "updateTime",
}

// NewVisitStatisticsDao creates and returns a new DAO object for table data access.
func NewVisitStatisticsDao() *VisitStatisticsDao {
	return &VisitStatisticsDao{
		group:   "default",
		table:   "visit_statistics",
		columns: visitStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VisitStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VisitStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VisitStatisticsDao) Columns() VisitStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VisitStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VisitStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VisitStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
