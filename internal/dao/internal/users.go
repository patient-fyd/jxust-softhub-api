// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UsersDao is the data access object for table users.
type UsersDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns UsersColumns // columns contains all the column names of Table for convenient usage.
}

// UsersColumns defines and stores column names for table users.
type UsersColumns struct {
	UserId     string // 用户ID，主键，自增
	UserName   string // 用户名，登录和显示名称
	Password   string // 用户密码，存储加密后的密码
	Name       string // 真实姓名
	RoleId     string // 角色ID，关联 roles 表，标识用户所属角色
	Avatar     string // 用户头像图片URL
	JoinYear   string // 入会年份，格式如2025
	Email      string // 用户邮箱
	Phone      string // 用户联系电话
	CreateTime string // 记录创建时间
	UpdateTime string // 记录最后更新时间
}

// usersColumns holds the columns for table users.
var usersColumns = UsersColumns{
	UserId:     "userId",
	UserName:   "userName",
	Password:   "password",
	Name:       "name",
	RoleId:     "roleId",
	Avatar:     "avatar",
	JoinYear:   "joinYear",
	Email:      "email",
	Phone:      "phone",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
}

// NewUsersDao creates and returns a new DAO object for table data access.
func NewUsersDao() *UsersDao {
	return &UsersDao{
		group:   "default",
		table:   "users",
		columns: usersColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UsersDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UsersDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UsersDao) Columns() UsersColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UsersDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UsersDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UsersDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
