// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ActivityRegistrationsDao is the data access object for table activity_registrations.
type ActivityRegistrationsDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns ActivityRegistrationsColumns // columns contains all the column names of Table for convenient usage.
}

// ActivityRegistrationsColumns defines and stores column names for table activity_registrations.
type ActivityRegistrationsColumns struct {
	RegistrationId string // 报名记录ID，主键，自增
	ActivityId     string // 活动ID，关联 activities 表
	UserId         string // 报名用户ID，关联 users 表，如为空表示未登录报名
	Name           string // 报名者姓名
	StudentId      string // 报名者学号
	Contact        string // 报名者联系方式，如电话或邮箱
	Status         string // 报名状态：0-待审核, 1-通过, 2-拒绝
	CreateTime     string // 记录创建时间
	UpdateTime     string // 记录最后更新时间
}

// activityRegistrationsColumns holds the columns for table activity_registrations.
var activityRegistrationsColumns = ActivityRegistrationsColumns{
	RegistrationId: "registrationId",
	ActivityId:     "activityId",
	UserId:         "userId",
	Name:           "name",
	StudentId:      "studentId",
	Contact:        "contact",
	Status:         "status",
	CreateTime:     "createTime",
	UpdateTime:     "updateTime",
}

// NewActivityRegistrationsDao creates and returns a new DAO object for table data access.
func NewActivityRegistrationsDao() *ActivityRegistrationsDao {
	return &ActivityRegistrationsDao{
		group:   "default",
		table:   "activity_registrations",
		columns: activityRegistrationsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ActivityRegistrationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ActivityRegistrationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ActivityRegistrationsDao) Columns() ActivityRegistrationsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ActivityRegistrationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ActivityRegistrationsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ActivityRegistrationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
