// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Projects is the golang structure for table projects.
type Projects struct {
	ProjectId   uint        `json:"project_id"  description:"项目ID，主键，自增"`
	Name        string      `json:"name"        description:"项目名称"`
	Description string      `json:"description" description:"项目描述信息"`
	TechStack   string      `json:"tech_stack"  description:"使用的技术栈，如 GoFrame2, Vue3 等"`
	CoverImage  string      `json:"cover_image" description:"项目封面图片URL"`
	Link        string      `json:"link"        description:"项目链接或跳转地址"`
	CreateTime  *gtime.Time `json:"create_time" description:"记录创建时间"`
	UpdateTime  *gtime.Time `json:"update_time" description:"记录最后更新时间"`
}
