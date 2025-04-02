// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TagsDao is the data access object for table tags.
type TagsDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns TagsColumns // columns contains all the column names of Table for convenient usage.
}

// TagsColumns defines and stores column names for table tags.
type TagsColumns struct {
	TagId      string // 标签ID，主键，自增
	TagName    string // 标签名称
	TagType    string // 标签类型，如news、project、resource等
	CreateTime string // 创建时间
}

// tagsColumns holds the columns for table tags.
var tagsColumns = TagsColumns{
	TagId:      "tagId",
	TagName:    "tagName",
	TagType:    "tagType",
	CreateTime: "createTime",
}

// NewTagsDao creates and returns a new DAO object for table data access.
func NewTagsDao() *TagsDao {
	return &TagsDao{
		group:   "default",
		table:   "tags",
		columns: tagsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TagsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TagsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TagsDao) Columns() TagsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TagsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TagsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TagsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
