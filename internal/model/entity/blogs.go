// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Blogs is the golang structure for table blogs.
type Blogs struct {
	BlogId       uint        `json:"blogId"       description:"博客ID，主键，自增"`
	Title        string      `json:"title"         description:"博客标题"`
	Content      string      `json:"content"       description:"博客内容，支持 Markdown 格式"`
	Summary      string      `json:"summary"       description:"博客摘要"`
	Category     string      `json:"category"      description:"博客分类，如前端、后端、移动开发、算法等"`
	Tags         string      `json:"tags"          description:"博客标签，多个标签用逗号分隔"`
	CoverImage   string      `json:"coverImage"   description:"封面图片的URL"`
	AuthorId     uint        `json:"authorId"     description:"作者ID，关联 users 表"`
	ViewCount    int         `json:"viewCount"    description:"浏览次数"`
	LikeCount    int         `json:"likeCount"    description:"点赞次数"`
	CommentCount int         `json:"commentCount" description:"评论次数"`
	IsRecommend  int         `json:"isRecommend"  description:"是否推荐：0-否，1-是"`
	Status       int         `json:"status"        description:"博客状态，0: 草稿, 1: 发布, 2: 下架"`
	CreateTime   *gtime.Time `json:"createTime"   description:"记录创建时间"`
	UpdateTime   *gtime.Time `json:"updateTime"   description:"记录最后更新时间"`
}
