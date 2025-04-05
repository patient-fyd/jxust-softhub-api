package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 获取评论列表请求
type ListReq struct {
	g.Meta      `path:"/api/comment/v1/list" method:"get" tags:"CommentService" summary:"获取评论列表"`
	ContentType string `p:"contentType" v:"required#内容类型不能为空" dc:"内容类型，例如：post"`
	ContentId   int    `p:"contentId" v:"required#内容ID不能为空" dc:"内容ID"`
	Page        int    `p:"page" v:"required|min:1#页码不能为空|页码必须大于0" dc:"页码"`
	Size        int    `p:"size" v:"required|min:1#每页条数不能为空|每页条数必须大于0" dc:"每页条数"`
}

// ListRes 获取评论列表响应
type ListRes struct {
	List  []CommentItem `json:"list"`  // 评论列表
	Total int           `json:"total"` // 总数
	Page  int           `json:"page"`  // 页码
	Size  int           `json:"size"`  // 每页条数
}

// CommentItem 评论列表项
type CommentItem struct {
	CommentId   int    `json:"commentId"`   // 评论ID
	ContentType string `json:"contentType"` // 内容类型
	ContentId   int    `json:"contentId"`   // 内容ID
	UserId      int    `json:"userId"`      // 用户ID
	UserName    string `json:"userName"`    // 用户名
	UserAvatar  string `json:"userAvatar"`  // 用户头像
	Content     string `json:"content"`     // 评论内容
	LikeCount   int    `json:"likeCount"`   // 点赞数
	IsLiked     bool   `json:"isLiked"`     // 当前用户是否已点赞
	CreateTime  string `json:"createTime"`  // 创建时间
}

// CreateReq 创建评论请求
type CreateReq struct {
	g.Meta      `path:"/api/comment/v1/create" method:"post" tags:"CommentService" summary:"创建评论"`
	ContentType string `p:"contentType" v:"required#内容类型不能为空" dc:"内容类型，例如：post"`
	ContentId   int    `p:"contentId" v:"required#内容ID不能为空" dc:"内容ID"`
	Content     string `p:"content" v:"required#评论内容不能为空" dc:"评论内容"`
}

// CreateRes 创建评论响应
type CreateRes struct {
	CommentId int `json:"commentId"` // 评论ID
}

// DeleteReq 删除评论请求
type DeleteReq struct {
	g.Meta    `path:"/api/comment/v1/delete" method:"post" tags:"CommentService" summary:"删除评论"`
	CommentId int `p:"commentId" v:"required#评论ID不能为空" dc:"评论ID"`
}

// DeleteRes 删除评论响应
type DeleteRes struct {
	Success bool `json:"success"` // 是否成功
}
