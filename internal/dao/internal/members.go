// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MembersDao is the data access object for table members.
type MembersDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns MembersColumns // columns contains all the column names of Table for convenient usage.
}

// MembersColumns defines and stores column names for table members.
type MembersColumns struct {
	MemberId     string // 成员ID，主键，自增
	UserId       string // 关联的用户ID
	Grade        string // 年级，如2020级
	JoinYear     string // 入会年份
	Department   string // 所属部门，如技术部、宣传部等
	Position     string // 职位，如部长、副部长、干事等
	Skills       string // 技能描述
	Introduction string // 个人简介
	IsCore       string // 是否为核心成员：0-否，1-是
	Status       string // 成员状态：0-待审核，1-正式成员，2-已退出
	CreateTime   string // 记录创建时间
	UpdateTime   string // 记录最后更新时间
}

// membersColumns holds the columns for table members.
var membersColumns = MembersColumns{
	MemberId:     "memberId",
	UserId:       "userId",
	Grade:        "grade",
	JoinYear:     "joinYear",
	Department:   "department",
	Position:     "position",
	Skills:       "skills",
	Introduction: "introduction",
	IsCore:       "isCore",
	Status:       "status",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
}

// NewMembersDao creates and returns a new DAO object for table data access.
func NewMembersDao() *MembersDao {
	return &MembersDao{
		group:   "default",
		table:   "members",
		columns: membersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MembersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MembersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MembersDao) Columns() MembersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MembersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MembersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MembersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
