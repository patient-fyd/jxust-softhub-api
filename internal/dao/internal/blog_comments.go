// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlogCommentsDao is the data access object for table blog_comments.
type BlogCommentsDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns BlogCommentsColumns // columns contains all the column names of Table for convenient usage.
}

// BlogCommentsColumns defines and stores column names for table blog_comments.
type BlogCommentsColumns struct {
	CommentId  string // 评论ID，主键，自增
	BlogId     string // 博客ID，关联blogs表
	UserId     string // 评论用户ID，关联users表
	Content    string // 评论内容
	ParentId   string // 父评论ID，用于回复功能，如为NULL则为顶级评论
	LikeCount  string // 点赞数
	Status     string // 状态：0-已删除，1-正常
	CreateTime string // 记录创建时间
	UpdateTime string // 记录最后更新时间
}

// blogCommentsColumns holds the columns for table blog_comments.
var blogCommentsColumns = BlogCommentsColumns{
	CommentId:  "commentId",
	BlogId:     "blogId",
	UserId:     "userId",
	Content:    "content",
	ParentId:   "parentId",
	LikeCount:  "likeCount",
	Status:     "status",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
}

// NewBlogCommentsDao creates and returns a new DAO object for table data access.
func NewBlogCommentsDao() *BlogCommentsDao {
	return &BlogCommentsDao{
		group:   "default",
		table:   "blog_comments",
		columns: blogCommentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlogCommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlogCommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlogCommentsDao) Columns() BlogCommentsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlogCommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlogCommentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlogCommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
