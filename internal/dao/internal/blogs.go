// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BlogsDao is the data access object for table blogs.
type BlogsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns BlogsColumns // columns contains all the column names of Table for convenient usage.
}

// BlogsColumns defines and stores column names for table blogs.
type BlogsColumns struct {
	BlogId       string // 博客ID，主键，自增
	Title        string // 博客标题
	Content      string // 博客内容，支持 Markdown 格式
	Summary      string // 博客摘要
	Category     string // 博客分类，如前端、后端、移动开发、算法等
	Tags         string // 博客标签，多个标签用逗号分隔
	CoverImage   string // 封面图片的URL
	AuthorId     string // 作者ID，关联 users 表
	ViewCount    string // 浏览次数
	LikeCount    string // 点赞次数
	CommentCount string // 评论次数
	IsRecommend  string // 是否推荐：0-否，1-是
	Status       string // 博客状态，0: 草稿, 1: 发布, 2: 下架
	CreateTime   string // 记录创建时间
	UpdateTime   string // 记录最后更新时间
}

// blogsColumns holds the columns for table blogs.
var blogsColumns = BlogsColumns{
	BlogId:       "blogId",
	Title:        "title",
	Content:      "content",
	Summary:      "summary",
	Category:     "category",
	Tags:         "tags",
	CoverImage:   "coverImage",
	AuthorId:     "authorId",
	ViewCount:    "viewCount",
	LikeCount:    "likeCount",
	CommentCount: "commentCount",
	IsRecommend:  "isRecommend",
	Status:       "status",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
}

// NewBlogsDao creates and returns a new DAO object for table data access.
func NewBlogsDao() *BlogsDao {
	return &BlogsDao{
		group:   "default",
		table:   "blogs",
		columns: blogsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *BlogsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *BlogsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *BlogsDao) Columns() BlogsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *BlogsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *BlogsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *BlogsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
