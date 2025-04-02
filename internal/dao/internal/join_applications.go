// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// JoinApplicationsDao is the data access object for table join_applications.
type JoinApplicationsDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns JoinApplicationsColumns // columns contains all the column names of Table for convenient usage.
}

// JoinApplicationsColumns defines and stores column names for table join_applications.
type JoinApplicationsColumns struct {
	ApplicationId    string // 申请ID，主键，自增
	Name             string // 申请人姓名
	StudentId        string // 学号
	Grade            string // 年级，如2020级
	College          string // 学院
	Major            string // 专业
	Phone            string // 联系电话
	Email            string // 邮箱
	Reason           string // 申请理由
	Skills           string // 技能介绍
	ExpectDepartment string // 期望加入的部门
	Status           string // 申请状态：0-待审核，1-通过，2-拒绝
	ReviewerId       string // 审核人ID，关联users表
	ReviewComment    string // 审核意见
	CreateTime       string // 申请时间
	UpdateTime       string // 最后更新时间
}

// joinApplicationsColumns holds the columns for table join_applications.
var joinApplicationsColumns = JoinApplicationsColumns{
	ApplicationId:    "applicationId",
	Name:             "name",
	StudentId:        "studentId",
	Grade:            "grade",
	College:          "college",
	Major:            "major",
	Phone:            "phone",
	Email:            "email",
	Reason:           "reason",
	Skills:           "skills",
	ExpectDepartment: "expectDepartment",
	Status:           "status",
	ReviewerId:       "reviewerId",
	ReviewComment:    "reviewComment",
	CreateTime:       "createTime",
	UpdateTime:       "updateTime",
}

// NewJoinApplicationsDao creates and returns a new DAO object for table data access.
func NewJoinApplicationsDao() *JoinApplicationsDao {
	return &JoinApplicationsDao{
		group:   "default",
		table:   "join_applications",
		columns: joinApplicationsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *JoinApplicationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *JoinApplicationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *JoinApplicationsDao) Columns() JoinApplicationsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *JoinApplicationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *JoinApplicationsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *JoinApplicationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
