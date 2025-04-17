// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlogLikesDao is the data access object for table blog_likes.
type BlogLikesDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns BlogLikesColumns // columns contains all the column names of Table for convenient usage.
}

// BlogLikesColumns defines and stores column names for table blog_likes.
type BlogLikesColumns struct {
	LikeId     string // 点赞ID，主键，自增
	BlogId     string // 博客ID，关联blogs表
	UserId     string // 用户ID，关联users表
	CreateTime string // 点赞时间
}

// blogLikesColumns holds the columns for table blog_likes.
var blogLikesColumns = BlogLikesColumns{
	LikeId:     "likeId",
	BlogId:     "blogId",
	UserId:     "userId",
	CreateTime: "createTime",
}

// NewBlogLikesDao creates and returns a new DAO object for table data access.
func NewBlogLikesDao() *BlogLikesDao {
	return &BlogLikesDao{
		group:   "default",
		table:   "blog_likes",
		columns: blogLikesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlogLikesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlogLikesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlogLikesDao) Columns() BlogLikesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlogLikesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlogLikesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlogLikesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
