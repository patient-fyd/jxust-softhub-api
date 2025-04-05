package v1

import (
	"github.com/gogf/gf/v2/frame/g"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// CommentListReq 获取评论列表请求
type CommentListReq struct {
	g.Meta      `path:"/list" method:"get" tags:"评论管理" summary:"获取评论列表"`
	ContentType string `json:"contentType" v:"required#内容类型不能为空"` // 内容类型: post:帖子
	ContentId   int    `json:"contentId" v:"required#内容ID不能为空"`   // 内容ID
	Page        int    `json:"page" d:"1"`                        // 页码
	Size        int    `json:"size" d:"10"`                       // 每页数量
}

// CommentListRes 获取评论列表响应
type CommentListRes struct {
	List  []model.CommentListItem `json:"list"`  // 评论列表
	Page  int                     `json:"page"`  // 页码
	Size  int                     `json:"size"`  // 每页数量
	Total int                     `json:"total"` // 总数
}

// CommentCreateReq 创建评论请求
type CommentCreateReq struct {
	g.Meta      `path:"/create" method:"post" tags:"评论管理" summary:"创建评论"`
	ContentType string `json:"contentType" v:"required#内容类型不能为空"` // 内容类型: post:帖子
	ContentId   int    `json:"contentId" v:"required#内容ID不能为空"`   // 内容ID
	Content     string `json:"content" v:"required#评论内容不能为空"`     // 评论内容
}

// CommentCreateRes 创建评论响应
type CommentCreateRes struct {
	CommentId int `json:"commentId"` // 评论ID
}

// CommentDeleteReq 删除评论请求
type CommentDeleteReq struct {
	g.Meta    `path:"/delete" method:"post" tags:"评论管理" summary:"删除评论"`
	CommentId int `json:"commentId" v:"required#评论ID不能为空"` // 评论ID
}

// CommentDeleteRes 删除评论响应
type CommentDeleteRes struct {
	Success bool `json:"success"` // 是否成功
}
