// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// News is the golang structure for table news.
type News struct {
	NewsId     uint        `json:"news_id"     description:"新闻ID，主键，自增"`
	Title      string      `json:"title"       description:"新闻标题"`
	Content    string      `json:"content"     description:"新闻内容，支持 Markdown 格式"`
	Category   string      `json:"category"    description:"新闻分类，如协会新闻、技术分享、赛事通知等"`
	NewsType   int         `json:"news_type"   description:"新闻类型：1-协会通知，2-技术分享"`
	CoverImage string      `json:"cover_image" description:"封面图片的URL"`
	AuthorId   uint        `json:"author_id"   description:"作者ID，关联 users 表"`
	ViewCount  int         `json:"view_count"  description:"浏览次数"`
	Status     int         `json:"status"      description:"新闻状态，0: 草稿, 1: 发布"`
	CreateTime *gtime.Time `json:"create_time" description:"记录创建时间"`
	UpdateTime *gtime.Time `json:"update_time" description:"记录最后更新时间"`
}
