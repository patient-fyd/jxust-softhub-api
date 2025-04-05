// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommentsDao is the data access object for table comments.
type CommentsDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CommentsColumns // columns contains all the column names of Table for convenient usage.
}

// CommentsColumns defines and stores column names for table comments.
type CommentsColumns struct {
	CommentId   string // 评论ID，主键，自增
	ContentType string // 评论关联内容类型，如 "news" 或 "activity"
	ContentId   string // 关联内容ID，指向具体的新闻或活动
	UserId      string // 评论用户ID，关联 users 表
	Content     string // 评论内容
	CreateTime  string // 记录创建时间
	LikeCount   string // 点赞数
}

// commentsColumns holds the columns for table comments.
var commentsColumns = CommentsColumns{
	CommentId:   "commentId",
	ContentType: "contentType",
	ContentId:   "contentId",
	UserId:      "userId",
	Content:     "content",
	CreateTime:  "createTime",
	LikeCount:   "likeCount",
}

// NewCommentsDao creates and returns a new DAO object for table data access.
func NewCommentsDao() *CommentsDao {
	return &CommentsDao{
		group:   "default",
		table:   "comments",
		columns: commentsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CommentsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CommentsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CommentsDao) Columns() CommentsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CommentsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CommentsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CommentsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
