// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivitiesDao is the data access object for table activities.
type ActivitiesDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns ActivitiesColumns // columns contains all the column names of Table for convenient usage.
}

// ActivitiesColumns defines and stores column names for table activities.
type ActivitiesColumns struct {
	ActivityId      string // 活动ID，主键，自增
	Title           string // 活动标题
	Description     string // 活动详细描述
	StartTime       string // 活动开始时间
	EndTime         string // 活动结束时间
	Location        string // 活动举办地点
	MaxParticipants string // 最大参与人数限制
	Status          string // 活动状态：0-未开始, 1-进行中, 2-已结束
	CreateTime      string // 记录创建时间
	UpdateTime      string // 记录最后更新时间
}

// activitiesColumns holds the columns for table activities.
var activitiesColumns = ActivitiesColumns{
	ActivityId:      "activityId",
	Title:           "title",
	Description:     "description",
	StartTime:       "startTime",
	EndTime:         "endTime",
	Location:        "location",
	MaxParticipants: "maxParticipants",
	Status:          "status",
	CreateTime:      "createTime",
	UpdateTime:      "updateTime",
}

// NewActivitiesDao creates and returns a new DAO object for table data access.
func NewActivitiesDao() *ActivitiesDao {
	return &ActivitiesDao{
		group:   "default",
		table:   "activities",
		columns: activitiesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivitiesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivitiesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivitiesDao) Columns() ActivitiesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivitiesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivitiesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivitiesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
