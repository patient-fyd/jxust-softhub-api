// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Messages is the golang structure for table messages.
type Messages struct {
	MessageId  uint        `json:"message_id"  description:"消息ID，主键，自增"`
	SenderId   uint        `json:"sender_id"   description:"发送者ID，关联 users 表"`
	ReceiverId uint        `json:"receiver_id" description:"接收者ID，关联 users 表"`
	Content    string      `json:"content"     description:"消息内容"`
	ReadStatus int         `json:"read_status" description:"读取状态：0-未读, 1-已读"`
	CreateTime *gtime.Time `json:"create_time" description:"记录发送时间"`
}
