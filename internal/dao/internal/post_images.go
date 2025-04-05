// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PostImagesDao is the data access object for table postImages.
type PostImagesDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns PostImagesColumns // columns contains all the column names of Table for convenient usage.
}

// PostImagesColumns defines and stores column names for table postImages.
type PostImagesColumns struct {
	ImageId    string // 图片ID
	PostId     string // 帖子ID
	ImageUrl   string // 图片URL
	SortOrder  string // 排序顺序
	CreateTime string // 创建时间
}

// postImagesColumns holds the columns for table postImages.
var postImagesColumns = PostImagesColumns{
	ImageId:    "imageId",
	PostId:     "postId",
	ImageUrl:   "imageUrl",
	SortOrder:  "sortOrder",
	CreateTime: "createTime",
}

// NewPostImagesDao creates and returns a new DAO object for table data access.
func NewPostImagesDao() *PostImagesDao {
	return &PostImagesDao{
		group:   "default",
		table:   "postImages",
		columns: postImagesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PostImagesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PostImagesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PostImagesDao) Columns() PostImagesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PostImagesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PostImagesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PostImagesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
