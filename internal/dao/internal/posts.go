// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PostsDao is the data access object for table posts.
type PostsDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns PostsColumns // columns contains all the column names of Table for convenient usage.
}

// PostsColumns defines and stores column names for table posts.
type PostsColumns struct {
	PostId       string // 帖子ID
	UserId       string // 发帖用户ID
	Content      string // 帖子内容
	CircleId     string // 所属圈子ID
	TopicId      string // 所属话题ID
	ViewCount    string // 浏览量
	LikeCount    string // 点赞数
	CommentCount string // 评论数
	ShareCount   string // 分享数
	IsTop        string // 是否置顶：0-否，1-是
	Status       string // 状态：0-草稿，1-已发布，2-已删除
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
}

// postsColumns holds the columns for table posts.
var postsColumns = PostsColumns{
	PostId:       "postId",
	UserId:       "userId",
	Content:      "content",
	CircleId:     "circleId",
	TopicId:      "topicId",
	ViewCount:    "viewCount",
	LikeCount:    "likeCount",
	CommentCount: "commentCount",
	ShareCount:   "shareCount",
	IsTop:        "isTop",
	Status:       "status",
	CreateTime:   "createTime",
	UpdateTime:   "updateTime",
}

// NewPostsDao creates and returns a new DAO object for table data access.
func NewPostsDao() *PostsDao {
	return &PostsDao{
		group:   "default",
		table:   "posts",
		columns: postsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PostsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PostsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PostsDao) Columns() PostsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PostsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PostsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PostsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
