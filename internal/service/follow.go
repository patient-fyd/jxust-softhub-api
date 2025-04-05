package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// IFollow 关注服务接口
type IFollow interface {
	// Toggle 关注/取消关注
	Toggle(ctx context.Context, in model.FollowToggleReq) (*model.FollowToggleOutput, error)
	// FollowingList 获取关注列表
	FollowingList(ctx context.Context, in model.FollowingListReq) (*model.FollowingListOutput, error)
	// FollowerList 获取粉丝列表
	FollowerList(ctx context.Context, in model.FollowerListReq) (*model.FollowerListOutput, error)
}

// 获取关注服务
func Follow() IFollow {
	return localFollow
}

// 注册关注服务
func RegisterFollow(i IFollow) {
	localFollow = i
}

var localFollow IFollow
