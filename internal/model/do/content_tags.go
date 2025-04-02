// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ContentTags is the golang structure of table content_tags for DAO operations like Where/Data.
type ContentTags struct {
	g.Meta      `orm:"table:content_tags, do:true"`
	ContentType interface{} // 内容类型，如news、project、resource等
	ContentId   interface{} // 内容ID
	TagId       interface{} // 标签ID，关联tags表
}
