package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LikeToggleReq 点赞/取消点赞请求
type LikeToggleReq struct {
	g.Meta     `path:"/like-toggle" method:"post" tags:"点赞管理" summary:"点赞/取消点赞"`
	TargetId   int `json:"targetId" v:"required#目标ID不能为空"`   // 目标ID
	TargetType int `json:"targetType" v:"required#目标类型不能为空"` // 目标类型：1-帖子，2-评论
}

// LikeToggleRes 点赞/取消点赞响应
type LikeToggleRes struct {
	LikeId   uint   `json:"likeId"`   // 点赞ID
	IsLiked  bool   `json:"isLiked"`  // 是否已点赞
	LikeTime string `json:"likeTime"` // 点赞时间
}
