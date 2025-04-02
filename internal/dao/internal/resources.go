// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ResourcesDao is the data access object for table resources.
type ResourcesDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns ResourcesColumns // columns contains all the column names of Table for convenient usage.
}

// ResourcesColumns defines and stores column names for table resources.
type ResourcesColumns struct {
	ResourceId  string // 资源ID，主键，自增
	Title       string // 资源标题
	Description string // 资源描述信息
	Category    string // 资源分类，如编程语言、算法、工具等
	FilePath    string // 资源文件存储路径或 URL
	UploaderId  string // 上传者ID，关联 users 表
	Downloads   string // 下载次数
	CreateTime  string // 记录创建时间
}

// resourcesColumns holds the columns for table resources.
var resourcesColumns = ResourcesColumns{
	ResourceId:  "resourceId",
	Title:       "title",
	Description: "description",
	Category:    "category",
	FilePath:    "filePath",
	UploaderId:  "uploaderId",
	Downloads:   "downloads",
	CreateTime:  "createTime",
}

// NewResourcesDao creates and returns a new DAO object for table data access.
func NewResourcesDao() *ResourcesDao {
	return &ResourcesDao{
		group:   "default",
		table:   "resources",
		columns: resourcesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ResourcesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ResourcesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ResourcesDao) Columns() ResourcesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ResourcesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ResourcesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ResourcesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
