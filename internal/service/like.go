package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// ILike 点赞服务接口
type ILike interface {
	// Toggle 点赞/取消点赞
	Toggle(ctx context.Context, in model.LikeToggleReq) (*model.LikeToggleOutput, error)
}

// 获取点赞服务
func Like() ILike {
	return localLike
}

// 注册点赞服务
func RegisterLike(i ILike) {
	localLike = i
}

var localLike ILike
