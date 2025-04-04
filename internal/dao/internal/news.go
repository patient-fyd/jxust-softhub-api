// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NewsDao is the data access object for table news.
type NewsDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns NewsColumns // columns contains all the column names of Table for convenient usage.
}

// NewsColumns defines and stores column names for table news.
type NewsColumns struct {
	NewsId     string // 新闻ID，主键，自增
	Title      string // 新闻标题
	Content    string // 新闻内容，支持 Markdown 格式
	Category   string // 新闻分类，如协会新闻、技术分享、赛事通知等
	NewsType   string // 新闻类型：1-协会通知，2-技术分享
	CoverImage string // 封面图片的URL
	AuthorId   string // 作者ID，关联 users 表
	ViewCount  string // 浏览次数
	Status     string // 新闻状态，0: 草稿, 1: 发布
	CreateTime string // 记录创建时间
	UpdateTime string // 记录最后更新时间
}

// newsColumns holds the columns for table news.
var newsColumns = NewsColumns{
	NewsId:     "newsId",
	Title:      "title",
	Content:    "content",
	Category:   "category",
	NewsType:   "newsType",
	CoverImage: "coverImage",
	AuthorId:   "authorId",
	ViewCount:  "viewCount",
	Status:     "status",
	CreateTime: "createTime",
	UpdateTime: "updateTime",
}

// NewNewsDao creates and returns a new DAO object for table data access.
func NewNewsDao() *NewsDao {
	return &NewsDao{
		group:   "default",
		table:   "news",
		columns: newsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NewsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NewsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NewsDao) Columns() NewsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NewsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NewsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NewsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
