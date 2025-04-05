// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LikesDao is the data access object for table likes.
type LikesDao struct {
	table   string       // table is the underlying table name of the DAO.
	group   string       // group is the database configuration group name of current DAO.
	columns LikesColumns // columns contains all the column names of Table for convenient usage.
}

// LikesColumns defines and stores column names for table likes.
type LikesColumns struct {
	LikeId     string // 点赞ID
	UserId     string // 用户ID
	TargetId   string // 目标ID
	TargetType string // 目标类型：1-帖子，2-评论
	CreateTime string // 创建时间
}

// likesColumns holds the columns for table likes.
var likesColumns = LikesColumns{
	LikeId:     "likeId",
	UserId:     "userId",
	TargetId:   "targetId",
	TargetType: "targetType",
	CreateTime: "createTime",
}

// NewLikesDao creates and returns a new DAO object for table data access.
func NewLikesDao() *LikesDao {
	return &LikesDao{
		group:   "default",
		table:   "likes",
		columns: likesColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LikesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LikesDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LikesDao) Columns() LikesColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LikesDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LikesDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LikesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
