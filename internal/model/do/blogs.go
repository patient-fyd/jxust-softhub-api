// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Blogs is the golang structure of table blogs for DAO operations like Where/Data.
type Blogs struct {
	g.Meta       `orm:"table:blogs, do:true"`
	BlogId       interface{} // 博客ID，主键，自增
	Title        interface{} // 博客标题
	Content      interface{} // 博客内容，支持 Markdown 格式
	Summary      interface{} // 博客摘要
	Category     interface{} // 博客分类，如前端、后端、移动开发、算法等
	Tags         interface{} // 博客标签，多个标签用逗号分隔
	CoverImage   interface{} // 封面图片的URL
	AuthorId     interface{} // 作者ID，关联 users 表
	ViewCount    interface{} // 浏览次数
	LikeCount    interface{} // 点赞次数
	CommentCount interface{} // 评论次数
	IsRecommend  interface{} // 是否推荐：0-否，1-是
	Status       interface{} // 博客状态，0: 草稿, 1: 发布, 2: 下架
	CreateTime   *gtime.Time // 记录创建时间
	UpdateTime   *gtime.Time // 记录最后更新时间
}
