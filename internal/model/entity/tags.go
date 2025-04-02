// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Tags is the golang structure for table tags.
type Tags struct {
	TagId      uint        `json:"tag_id"      description:"标签ID，主键，自增"`
	TagName    string      `json:"tag_name"    description:"标签名称"`
	TagType    string      `json:"tag_type"    description:"标签类型，如news、project、resource等"`
	CreateTime *gtime.Time `json:"create_time" description:"创建时间"`
}
