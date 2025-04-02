// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Tags is the golang structure of table tags for DAO operations like Where/Data.
type Tags struct {
	g.Meta     `orm:"table:tags, do:true"`
	TagId      interface{} // 标签ID，主键，自增
	TagName    interface{} // 标签名称
	TagType    interface{} // 标签类型，如news、project、resource等
	CreateTime *gtime.Time // 创建时间
}
