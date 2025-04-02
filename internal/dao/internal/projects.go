// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ProjectsDao is the data access object for table projects.
type ProjectsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ProjectsColumns // columns contains all the column names of Table for convenient usage.
}

// ProjectsColumns defines and stores column names for table projects.
type ProjectsColumns struct {
	ProjectId   string // 项目ID，主键，自增
	Name        string // 项目名称
	Description string // 项目描述信息
	TechStack   string // 使用的技术栈，如 GoFrame2, Vue3 等
	CoverImage  string // 项目封面图片URL
	Link        string // 项目链接或跳转地址
	CreateTime  string // 记录创建时间
	UpdateTime  string // 记录最后更新时间
}

// projectsColumns holds the columns for table projects.
var projectsColumns = ProjectsColumns{
	ProjectId:   "projectId",
	Name:        "name",
	Description: "description",
	TechStack:   "techStack",
	CoverImage:  "coverImage",
	Link:        "link",
	CreateTime:  "createTime",
	UpdateTime:  "updateTime",
}

// NewProjectsDao creates and returns a new DAO object for table data access.
func NewProjectsDao() *ProjectsDao {
	return &ProjectsDao{
		group:   "default",
		table:   "projects",
		columns: projectsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ProjectsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ProjectsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ProjectsDao) Columns() ProjectsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ProjectsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ProjectsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ProjectsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
