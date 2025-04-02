// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotificationsDao is the data access object for table notifications.
type NotificationsDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns NotificationsColumns // columns contains all the column names of Table for convenient usage.
}

// NotificationsColumns defines and stores column names for table notifications.
type NotificationsColumns struct {
	NotificationId string // 通知ID，主键，自增
	UserId         string // 接收通知的用户ID，关联 users 表
	Title          string // 通知标题
	Content        string // 通知内容
	Type           string // 通知类型，如系统通知、活动提醒等
	ReadStatus     string // 读取状态：0-未读, 1-已读
	CreateTime     string // 记录发送时间
}

// notificationsColumns holds the columns for table notifications.
var notificationsColumns = NotificationsColumns{
	NotificationId: "notificationId",
	UserId:         "userId",
	Title:          "title",
	Content:        "content",
	Type:           "type",
	ReadStatus:     "readStatus",
	CreateTime:     "createTime",
}

// NewNotificationsDao creates and returns a new DAO object for table data access.
func NewNotificationsDao() *NotificationsDao {
	return &NotificationsDao{
		group:   "default",
		table:   "notifications",
		columns: notificationsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NotificationsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NotificationsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NotificationsDao) Columns() NotificationsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NotificationsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NotificationsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NotificationsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
