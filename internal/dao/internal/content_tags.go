// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ContentTagsDao is the data access object for table content_tags.
type ContentTagsDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ContentTagsColumns // columns contains all the column names of Table for convenient usage.
}

// ContentTagsColumns defines and stores column names for table content_tags.
type ContentTagsColumns struct {
	ContentType string // 内容类型，如news、project、resource等
	ContentId   string // 内容ID
	TagId       string // 标签ID，关联tags表
}

// contentTagsColumns holds the columns for table content_tags.
var contentTagsColumns = ContentTagsColumns{
	ContentType: "contentType",
	ContentId:   "contentId",
	TagId:       "tagId",
}

// NewContentTagsDao creates and returns a new DAO object for table data access.
func NewContentTagsDao() *ContentTagsDao {
	return &ContentTagsDao{
		group:   "default",
		table:   "content_tags",
		columns: contentTagsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ContentTagsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ContentTagsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ContentTagsDao) Columns() ContentTagsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ContentTagsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ContentTagsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ContentTagsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
