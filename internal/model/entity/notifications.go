// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Notifications is the golang structure for table notifications.
type Notifications struct {
	NotificationId uint        `json:"notification_id" description:"通知ID，主键，自增"`
	UserId         uint        `json:"user_id"         description:"接收通知的用户ID，关联 users 表"`
	Title          string      `json:"title"           description:"通知标题"`
	Content        string      `json:"content"         description:"通知内容"`
	Type           string      `json:"type"            description:"通知类型，如系统通知、活动提醒等"`
	ReadStatus     int         `json:"read_status"     description:"读取状态：0-未读, 1-已读"`
	CreateTime     *gtime.Time `json:"create_time"     description:"记录发送时间"`
}
