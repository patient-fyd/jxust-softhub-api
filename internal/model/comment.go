package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentListInput 评论列表输入
type CommentListInput struct {
	ContentType string `json:"contentType" v:"required#内容类型不能为空"` // 内容类型: post:帖子
	ContentId   int    `json:"contentId" v:"required#内容ID不能为空"`   // 内容ID
	Page        int    `json:"page"`                              // 页码
	Size        int    `json:"size"`                              // 每页数量
}

// CommentListItem 评论列表项
type CommentListItem struct {
	CommentId   int         `json:"commentId"`   // 评论ID
	ContentType string      `json:"contentType"` // 内容类型
	ContentId   int         `json:"contentId"`   // 内容ID
	UserId      int         `json:"userId"`      // 用户ID
	UserName    string      `json:"userName"`    // 用户名
	UserAvatar  string      `json:"userAvatar"`  // 用户头像
	Content     string      `json:"content"`     // 评论内容
	LikeCount   int         `json:"likeCount"`   // 点赞数
	IsLiked     bool        `json:"isLiked"`     // 当前用户是否已点赞
	CreateTime  *gtime.Time `json:"createTime"`  // 创建时间
}

// CommentListOutput 评论列表输出
type CommentListOutput struct {
	List  []CommentListItem `json:"list"`  // 评论列表
	Page  int               `json:"page"`  // 页码
	Size  int               `json:"size"`  // 每页数量
	Total int               `json:"total"` // 总数
}

// CommentCreateInput 创建评论输入
type CommentCreateInput struct {
	ContentType string `json:"contentType" v:"required#内容类型不能为空"` // 内容类型: post:帖子
	ContentId   int    `json:"contentId" v:"required#内容ID不能为空"`   // 内容ID
	Content     string `json:"content" v:"required#评论内容不能为空"`     // 评论内容
}

// CommentCreateOutput 创建评论输出
type CommentCreateOutput struct {
	CommentId int `json:"commentId"` // 评论ID
}

// CommentDeleteInput 删除评论输入
type CommentDeleteInput struct {
	CommentId int `json:"commentId" v:"required#评论ID不能为空"` // 评论ID
}

// CommentDeleteOutput 删除评论输出
type CommentDeleteOutput struct {
	Success bool `json:"success"` // 是否成功
}
