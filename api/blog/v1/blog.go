package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// BlogListReq 获取博客列表请求
type BlogListReq struct {
	g.Meta      `path:"/api/blog/v1/blog/list" method:"get" tags:"BlogService" summary:"获取博客列表"`
	Page        int    `p:"page" v:"required|min:1#页码不能为空|页码必须大于0" d:"1" dc:"页码"`
	Size        int    `p:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0" d:"10" dc:"每页条数"`
	Category    string `p:"category" dc:"博客分类，可选"`
	Tag         string `p:"tag" dc:"博客标签，可选"`
	Keyword     string `p:"keyword" dc:"关键词搜索(标题、摘要)，可选"`
	AuthorId    uint   `p:"authorId" dc:"作者ID，可选"`
	Status      int    `p:"status" d:"1" dc:"状态: 0-草稿，1-发布，2-下架，默认为1"`
	IsRecommend int    `p:"isRecommend" d:"0" dc:"是否推荐: 0-全部，1-仅推荐，默认为0"`
}

// BlogListRes 获取博客列表响应
type BlogListRes struct {
	List  []*model.BlogListItem `json:"list"`  // 博客列表
	Page  int                   `json:"page"`  // 页码
	Size  int                   `json:"size"`  // 每页记录数
	Total int                   `json:"total"` // 总记录数
}

// BlogDetailReq 获取博客详情请求
type BlogDetailReq struct {
	g.Meta `path:"/api/blog/v1/blog/detail" method:"get" tags:"BlogService" summary:"获取博客详情"`
	BlogId uint `p:"blogId" v:"required#博客ID不能为空" dc:"博客ID"`
}

// BlogDetailRes 获取博客详情响应
type BlogDetailRes struct {
	*model.BlogDetailOutput
}

// BlogCreateReq 创建博客请求
type BlogCreateReq struct {
	g.Meta     `path:"/api/blog/v1/blog/create" method:"post" tags:"BlogService" summary:"创建博客"`
	Title      string `p:"title" v:"required#博客标题不能为空" dc:"博客标题"`
	Content    string `p:"content" v:"required#博客内容不能为空" dc:"博客内容，支持Markdown"`
	Summary    string `p:"summary" dc:"博客摘要，可选，不传则自动截取内容前部分"`
	Category   string `p:"category" v:"required#博客分类不能为空" dc:"博客分类"`
	Tags       string `p:"tags" dc:"博客标签，多个标签用逗号分隔，可选"`
	CoverImage string `p:"coverImage" dc:"封面图片URL，可选"`
	Status     int    `p:"status" v:"required|in:0,1,2#状态不能为空|状态只能为0,1,2" dc:"状态：0-草稿，1-发布，2-下架"`
}

// BlogCreateRes 创建博客响应
type BlogCreateRes struct {
	BlogId uint `json:"blogId"` // 博客ID
}

// BlogUpdateReq 更新博客请求
type BlogUpdateReq struct {
	g.Meta     `path:"/api/blog/v1/blog/update" method:"post" tags:"BlogService" summary:"更新博客"`
	BlogId     uint   `p:"blogId" v:"required#博客ID不能为空" dc:"博客ID"`
	Title      string `p:"title" v:"required#博客标题不能为空" dc:"博客标题"`
	Content    string `p:"content" v:"required#博客内容不能为空" dc:"博客内容，支持Markdown"`
	Summary    string `p:"summary" dc:"博客摘要，可选，不传则自动截取内容前部分"`
	Category   string `p:"category" v:"required#博客分类不能为空" dc:"博客分类"`
	Tags       string `p:"tags" dc:"博客标签，多个标签用逗号分隔，可选"`
	CoverImage string `p:"coverImage" dc:"封面图片URL，可选"`
	Status     int    `p:"status" v:"required|in:0,1,2#状态不能为空|状态只能为0,1,2" dc:"状态：0-草稿，1-发布，2-下架"`
}

// BlogUpdateRes 更新博客响应
type BlogUpdateRes struct {
	Success bool `json:"success"` // 是否成功
}

// BlogDeleteReq 删除博客请求
type BlogDeleteReq struct {
	g.Meta `path:"/api/blog/v1/blog/delete" method:"post" tags:"BlogService" summary:"删除博客"`
	BlogId uint `p:"blogId" v:"required#博客ID不能为空" dc:"博客ID"`
}

// BlogDeleteRes 删除博客响应
type BlogDeleteRes struct {
	Success bool `json:"success"` // 是否成功
}

// BlogRecommendReq 设置博客推荐状态请求
type BlogRecommendReq struct {
	g.Meta      `path:"/api/blog/v1/blog/recommend" method:"post" tags:"BlogService" summary:"设置博客推荐状态"`
	BlogId      uint `p:"blogId" v:"required#博客ID不能为空" dc:"博客ID"`
	IsRecommend int  `p:"isRecommend" v:"required|in:0,1#是否推荐不能为空|是否推荐只能为0,1" dc:"是否推荐：0-否，1-是"`
}

// BlogRecommendRes 设置博客推荐状态响应
type BlogRecommendRes struct {
	Success bool `json:"success"` // 是否成功
}

// BlogLikeReq 点赞博客请求
type BlogLikeReq struct {
	g.Meta `path:"/api/blog/v1/blog/like" method:"post" tags:"BlogService" summary:"点赞博客"`
	BlogId uint `p:"blogId" v:"required#博客ID不能为空" dc:"博客ID"`
}

// BlogLikeRes 点赞博客响应
type BlogLikeRes struct {
	Success bool `json:"success"` // 是否成功
}

// BlogUnlikeReq 取消点赞博客请求
type BlogUnlikeReq struct {
	g.Meta `path:"/api/blog/v1/blog/unlike" method:"post" tags:"BlogService" summary:"取消点赞博客"`
	BlogId uint `p:"blogId" v:"required#博客ID不能为空" dc:"博客ID"`
}

// BlogUnlikeRes 取消点赞博客响应
type BlogUnlikeRes struct {
	Success bool `json:"success"` // 是否成功
}

// BlogCommentListReq 获取博客评论列表请求
type BlogCommentListReq struct {
	g.Meta `path:"/api/blog/v1/blog/comment/list" method:"get" tags:"BlogService" summary:"获取博客评论列表"`
	BlogId uint `p:"blogId" v:"required#博客ID不能为空" dc:"博客ID"`
	Page   int  `p:"page" v:"min:1#页码最小值为1" d:"1" dc:"页码，默认为1"`
	Size   int  `p:"size" v:"max:50#每页最大记录数为50" d:"10" dc:"每页记录数，默认为10，最大为50"`
}

// BlogCommentListRes 获取博客评论列表响应
type BlogCommentListRes struct {
	List  []*model.BlogCommentItem `json:"list"`  // 评论列表
	Page  int                      `json:"page"`  // 页码
	Size  int                      `json:"size"`  // 每页记录数
	Total int                      `json:"total"` // 总记录数
}

// BlogCommentCreateReq 创建博客评论请求
type BlogCommentCreateReq struct {
	g.Meta   `path:"/api/blog/v1/blog/comment/create" method:"post" tags:"BlogService" summary:"创建博客评论"`
	BlogId   uint   `p:"blogId" v:"required#博客ID不能为空" dc:"博客ID"`
	Content  string `p:"content" v:"required#评论内容不能为空" dc:"评论内容"`
	ParentId uint   `p:"parentId" dc:"父评论ID，用于回复功能，可选"`
}

// BlogCommentCreateRes 创建博客评论响应
type BlogCommentCreateRes struct {
	CommentId uint `json:"commentId"` // 评论ID
}

// BlogCommentDeleteReq 删除博客评论请求
type BlogCommentDeleteReq struct {
	g.Meta    `path:"/api/blog/v1/blog/comment/delete" method:"post" tags:"BlogService" summary:"删除博客评论"`
	CommentId uint `p:"commentId" v:"required#评论ID不能为空" dc:"评论ID"`
}

// BlogCommentDeleteRes 删除博客评论响应
type BlogCommentDeleteRes struct {
	Success bool `json:"success"` // 是否成功
}

// BlogCategoryListReq 获取博客分类列表请求
type BlogCategoryListReq struct {
	g.Meta `path:"/api/blog/v1/blog/category/list" method:"get" tags:"BlogService" summary:"获取博客分类列表"`
}

// BlogCategoryListRes 获取博客分类列表响应
type BlogCategoryListRes struct {
	List []string `json:"list"` // 分类列表
}

// BlogTagListReq 获取博客标签列表请求
type BlogTagListReq struct {
	g.Meta `path:"/api/blog/v1/blog/tag/list" method:"get" tags:"BlogService" summary:"获取博客标签列表"`
}

// BlogTagListRes 获取博客标签列表响应
type BlogTagListRes struct {
	List []string `json:"list"` // 标签列表
}
