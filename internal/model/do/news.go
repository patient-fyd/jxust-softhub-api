// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// News is the golang structure of table news for DAO operations like Where/Data.
type News struct {
	g.Meta     `orm:"table:news, do:true"`
	NewsId     interface{} // 新闻ID，主键，自增
	Title      interface{} // 新闻标题
	Content    interface{} // 新闻内容，支持 Markdown 格式
	Category   interface{} // 新闻分类，如协会新闻、技术分享、赛事通知等
	NewsType   interface{} // 新闻类型：1-协会通知，2-技术分享
	CoverImage interface{} // 封面图片的URL
	AuthorId   interface{} // 作者ID，关联 users 表
	ViewCount  interface{} // 浏览次数
	Status     interface{} // 新闻状态，0: 草稿, 1: 发布
	CreateTime *gtime.Time // 记录创建时间
	UpdateTime *gtime.Time // 记录最后更新时间
}
