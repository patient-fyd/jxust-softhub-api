// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PostImages is the golang structure for table postImages.
type PostImages struct {
	ImageId    uint        `json:"image_id"    description:"图片ID"`
	PostId     uint        `json:"post_id"     description:"帖子ID"`
	ImageUrl   string      `json:"image_url"   description:"图片URL"`
	SortOrder  uint        `json:"sort_order"  description:"排序顺序"`
	CreateTime *gtime.Time `json:"create_time" description:"创建时间"`
}
