// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TopicsDao is the data access object for table topics.
type TopicsDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns TopicsColumns // columns contains all the column names of Table for convenient usage.
}

// TopicsColumns defines and stores column names for table topics.
type TopicsColumns struct {
	TopicId     string // 话题ID
	Name        string // 话题名称
	Description string // 话题描述
	Icon        string // 话题图标
	PostCount   string // 帖子数
	Status      string // 状态：0-已删除，1-正常
	IsHot       string // 是否热门：0-否，1-是
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
}

// topicsColumns holds the columns for table topics.
var topicsColumns = TopicsColumns{
	TopicId:     "topicId",
	Name:        "name",
	Description: "description",
	Icon:        "icon",
	PostCount:   "postCount",
	Status:      "status",
	IsHot:       "isHot",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
}

// NewTopicsDao creates and returns a new DAO object for table data access.
func NewTopicsDao() *TopicsDao {
	return &TopicsDao{
		group:   "default",
		table:   "topics",
		columns: topicsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TopicsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TopicsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TopicsDao) Columns() TopicsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TopicsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TopicsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TopicsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
