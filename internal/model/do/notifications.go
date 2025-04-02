// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notifications is the golang structure of table notifications for DAO operations like Where/Data.
type Notifications struct {
	g.Meta         `orm:"table:notifications, do:true"`
	NotificationId interface{} // 通知ID，主键，自增
	UserId         interface{} // 接收通知的用户ID，关联 users 表
	Title          interface{} // 通知标题
	Content        interface{} // 通知内容
	Type           interface{} // 通知类型，如系统通知、活动提醒等
	ReadStatus     interface{} // 读取状态：0-未读, 1-已读
	CreateTime     *gtime.Time // 记录发送时间
}
