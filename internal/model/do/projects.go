// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Projects is the golang structure of table projects for DAO operations like Where/Data.
type Projects struct {
	g.Meta      `orm:"table:projects, do:true"`
	ProjectId   interface{} // 项目ID，主键，自增
	Name        interface{} // 项目名称
	Description interface{} // 项目描述信息
	TechStack   interface{} // 使用的技术栈，如 GoFrame2, Vue3 等
	CoverImage  interface{} // 项目封面图片URL
	Link        interface{} // 项目链接或跳转地址
	CreateTime  *gtime.Time // 记录创建时间
	UpdateTime  *gtime.Time // 记录最后更新时间
}
