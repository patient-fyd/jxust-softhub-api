package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LikeToggleReq 点赞/取消点赞输入
type LikeToggleReq struct {
	TargetId   int `json:"targetId" v:"required#目标ID不能为空"`   // 目标ID
	TargetType int `json:"targetType" v:"required#目标类型不能为空"` // 目标类型：1-帖子，2-评论
}

// LikeToggleOutput 点赞/取消点赞输出
type LikeToggleOutput struct {
	LikeId   uint        `json:"likeId"`   // 点赞ID
	IsLiked  bool        `json:"isLiked"`  // 是否已点赞
	LikeTime *gtime.Time `json:"likeTime"` // 点赞时间
}
