package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
)

// 博客列表查询参数
type BlogListInput struct {
	Page        int    `json:"page" d:"1"`         // 页码，默认为1
	Size        int    `json:"size" d:"10"`        // 每页记录数，默认为10
	Category    string `json:"category"`           // 博客分类
	Tag         string `json:"tag"`                // 博客标签
	Keyword     string `json:"keyword"`            // 关键词搜索（标题、摘要）
	AuthorId    uint   `json:"author_id"`          // 作者ID
	Status      int    `json:"status" d:"1"`       // 状态: 0-草稿，1-发布，2-下架
	IsRecommend int    `json:"is_recommend" d:"0"` // 是否推荐: 0-全部，1-仅推荐
}

// 博客列表输出
type BlogListOutput struct {
	List  []*BlogListItem `json:"list"`  // 博客列表
	Page  int             `json:"page"`  // 页码
	Size  int             `json:"size"`  // 每页记录数
	Total int             `json:"total"` // 总记录数
}

// 博客列表项
type BlogListItem struct {
	BlogId       uint        `json:"blog_id"`       // 博客ID
	Title        string      `json:"title"`         // 博客标题
	Summary      string      `json:"summary"`       // 博客摘要
	Category     string      `json:"category"`      // 博客分类
	Tags         string      `json:"tags"`          // 博客标签
	CoverImage   string      `json:"cover_image"`   // 封面图片URL
	AuthorId     uint        `json:"author_id"`     // 作者ID
	AuthorName   string      `json:"author_name"`   // 作者昵称
	ViewCount    int         `json:"view_count"`    // 浏览次数
	LikeCount    int         `json:"like_count"`    // 点赞次数
	CommentCount int         `json:"comment_count"` // 评论次数
	IsRecommend  int         `json:"is_recommend"`  // 是否推荐
	Status       int         `json:"status"`        // 状态
	CreateTime   *gtime.Time `json:"create_time"`   // 创建时间
}

// 博客详情输入
type BlogDetailInput struct {
	BlogId uint `json:"blog_id"` // 博客ID
}

// 博客详情输出
type BlogDetailOutput struct {
	entity.Blogs        // 博客基本信息
	AuthorName   string `json:"author_name"`   // 作者昵称
	AuthorAvatar string `json:"author_avatar"` // 作者头像
	IsLiked      bool   `json:"is_liked"`      // 当前用户是否已点赞
}

// 博客创建输入
type BlogCreateInput struct {
	Title      string `json:"title" v:"required#博客标题不能为空"`                    // 博客标题
	Content    string `json:"content" v:"required#博客内容不能为空"`                  // 博客内容
	Summary    string `json:"summary"`                                        // 博客摘要
	Category   string `json:"category" v:"required#博客分类不能为空"`                 // 博客分类
	Tags       string `json:"tags"`                                           // 博客标签
	CoverImage string `json:"cover_image"`                                    // 封面图片URL
	Status     int    `json:"status" v:"required|in:0,1,2#状态不能为空|状态只能为0,1,2"` // 状态: 0-草稿，1-发布，2-下架
}

// 博客创建输出
type BlogCreateOutput struct {
	BlogId uint `json:"blog_id"` // 博客ID
}

// 博客更新输入
type BlogUpdateInput struct {
	BlogId     uint   `json:"blog_id" v:"required#博客ID不能为空"`                  // 博客ID
	Title      string `json:"title" v:"required#博客标题不能为空"`                    // 博客标题
	Content    string `json:"content" v:"required#博客内容不能为空"`                  // 博客内容
	Summary    string `json:"summary"`                                        // 博客摘要
	Category   string `json:"category" v:"required#博客分类不能为空"`                 // 博客分类
	Tags       string `json:"tags"`                                           // 博客标签
	CoverImage string `json:"cover_image"`                                    // 封面图片URL
	Status     int    `json:"status" v:"required|in:0,1,2#状态不能为空|状态只能为0,1,2"` // 状态: 0-草稿，1-发布，2-下架
}

// 博客删除输入
type BlogDeleteInput struct {
	BlogId uint `json:"blog_id" v:"required#博客ID不能为空"` // 博客ID
}

// 博客推荐设置输入
type BlogRecommendInput struct {
	BlogId      uint `json:"blog_id" v:"required#博客ID不能为空"`                        // 博客ID
	IsRecommend int  `json:"is_recommend" v:"required|in:0,1#是否推荐不能为空|是否推荐只能为0,1"` // 是否推荐: 0-否，1-是
}

// 博客点赞输入
type BlogLikeInput struct {
	BlogId uint `json:"blog_id" v:"required#博客ID不能为空"` // 博客ID
}

// 博客取消点赞输入
type BlogUnlikeInput struct {
	BlogId uint `json:"blog_id" v:"required#博客ID不能为空"` // 博客ID
}

// 博客评论列表输入
type BlogCommentListInput struct {
	BlogId uint `json:"blog_id" v:"required#博客ID不能为空"` // 博客ID
	Page   int  `json:"page" d:"1"`                    // 页码，默认为1
	Size   int  `json:"size" d:"10"`                   // 每页记录数，默认为10
}

// 博客评论列表输出
type BlogCommentListOutput struct {
	List  []*BlogCommentItem `json:"list"`  // 评论列表
	Page  int                `json:"page"`  // 页码
	Size  int                `json:"size"`  // 每页记录数
	Total int                `json:"total"` // 总记录数
}

// 博客评论项
type BlogCommentItem struct {
	CommentId  uint        `json:"comment_id"`  // 评论ID
	BlogId     uint        `json:"blog_id"`     // 博客ID
	UserId     uint        `json:"user_id"`     // 用户ID
	UserName   string      `json:"user_name"`   // 用户昵称
	UserAvatar string      `json:"user_avatar"` // 用户头像
	Content    string      `json:"content"`     // 评论内容
	ParentId   uint        `json:"parent_id"`   // 父评论ID
	LikeCount  uint        `json:"like_count"`  // 点赞数
	ReplyCount int         `json:"reply_count"` // 回复数
	CreateTime *gtime.Time `json:"create_time"` // 创建时间
	// 如果有父评论，则包含父评论信息
	ParentUserId   uint   `json:"parent_user_id,omitempty"`   // 父评论用户ID
	ParentUserName string `json:"parent_user_name,omitempty"` // 父评论用户昵称
}

// 博客评论创建输入
type BlogCommentCreateInput struct {
	BlogId   uint   `json:"blog_id" v:"required#博客ID不能为空"` // 博客ID
	Content  string `json:"content" v:"required#评论内容不能为空"` // 评论内容
	ParentId uint   `json:"parent_id"`                     // 父评论ID，可选
}

// 博客评论创建输出
type BlogCommentCreateOutput struct {
	CommentId uint `json:"comment_id"` // 评论ID
}

// 博客评论删除输入
type BlogCommentDeleteInput struct {
	CommentId uint `json:"comment_id" v:"required#评论ID不能为空"` // 评论ID
}

// 博客分类列表输出
type BlogCategoryListOutput struct {
	List []string `json:"list"` // 分类列表
}

// 博客标签列表输出
type BlogTagListOutput struct {
	List []string `json:"list"` // 标签列表
}
