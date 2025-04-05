package model

// LikeToggleInput 点赞/取消点赞输入参数
type LikeToggleInput struct {
	TargetId   int `json:"targetId" v:"required#目标ID不能为空"`   // 目标ID
	TargetType int `json:"targetType" v:"required#目标类型不能为空"` // 目标类型：1-帖子，2-评论
}

// LikeToggleOutput 点赞/取消点赞输出参数
type LikeToggleOutput struct {
	Status    int `json:"status"`    // 操作后状态：1-已点赞，0-已取消点赞
	LikeCount int `json:"likeCount"` // 点赞总数
}
