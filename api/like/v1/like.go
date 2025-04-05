package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ToggleReq 点赞/取消点赞请求
type ToggleReq struct {
	g.Meta     `path:"/api/like/v1/toggle" method:"post" tags:"LikeService" summary:"点赞/取消点赞"`
	TargetId   int `p:"targetId" v:"required#目标ID不能为空" dc:"目标ID"`
	TargetType int `p:"targetType" v:"required#目标类型不能为空" dc:"目标类型：1-帖子，2-评论"`
}

// ToggleRes 点赞/取消点赞响应
type ToggleRes struct {
	Status    int `json:"status"`    // 操作后状态：1-已点赞，0-已取消
	LikeCount int `json:"likeCount"` // 点赞总数
}
