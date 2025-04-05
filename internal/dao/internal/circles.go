// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CirclesDao is the data access object for table circles.
type CirclesDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns CirclesColumns // columns contains all the column names of Table for convenient usage.
}

// CirclesColumns defines and stores column names for table circles.
type CirclesColumns struct {
	CircleId    string // 圈子ID
	Name        string // 圈子名称
	Description string // 圈子描述
	Icon        string // 圈子图标
	UserId      string // 创建者ID
	PostCount   string // 帖子数
	MemberCount string // 成员数
	Status      string // 状态：0-已删除，1-正常
	IsOfficial  string // 是否官方圈子：0-否，1-是
	CreateTime  string // 创建时间
	UpdateTime  string // 更新时间
}

// circlesColumns holds the columns for table circles.
var circlesColumns = CirclesColumns{
	CircleId:    "circleId",
	Name:        "name",
	Description: "description",
	Icon:        "icon",
	UserId:      "userId",
	PostCount:   "postCount",
	MemberCount: "memberCount",
	Status:      "status",
	IsOfficial:  "isOfficial",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
}

// NewCirclesDao creates and returns a new DAO object for table data access.
func NewCirclesDao() *CirclesDao {
	return &CirclesDao{
		group:   "default",
		table:   "circles",
		columns: circlesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CirclesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CirclesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CirclesDao) Columns() CirclesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CirclesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CirclesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CirclesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
