// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PostImages is the golang structure of table postImages for DAO operations like Where/Data.
type PostImages struct {
	g.Meta     `orm:"table:postImages, do:true"`
	ImageId    interface{} // 图片ID
	PostId     interface{} // 帖子ID
	ImageUrl   interface{} // 图片URL
	SortOrder  interface{} // 排序顺序
	CreateTime *gtime.Time // 创建时间
}
