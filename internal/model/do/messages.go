// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Messages is the golang structure of table messages for DAO operations like Where/Data.
type Messages struct {
	g.Meta     `orm:"table:messages, do:true"`
	MessageId  interface{} // 消息ID，主键，自增
	SenderId   interface{} // 发送者ID，关联 users 表
	ReceiverId interface{} // 接收者ID，关联 users 表
	Content    interface{} // 消息内容
	ReadStatus interface{} // 读取状态：0-未读, 1-已读
	CreateTime *gtime.Time // 记录发送时间
}
