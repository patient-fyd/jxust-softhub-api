package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// ICircle 圈子服务接口
type ICircle interface {
	// List 获取圈子列表
	List(ctx context.Context, in model.CircleListInput) (*model.CircleListOutput, error)
	// Detail 获取圈子详情
	Detail(ctx context.Context, in model.CircleDetailInput) (*model.CircleDetailOutput, error)
	// Join 关注/取消关注圈子
	Join(ctx context.Context, in model.CircleJoinInput) (*model.CircleJoinOutput, error)
	// MyCircles 获取我关注的圈子列表
	MyCircles(ctx context.Context, in model.CircleMyCirclesInput) (*model.CircleMyCirclesOutput, error)
	// CircleStat 获取圈子统计信息
	CircleStat(ctx context.Context, in model.CircleStatInput) (*model.CircleStatOutput, error)
}

// 获取圈子服务
func Circle() ICircle {
	return localCircle
}

// 注册圈子服务
func RegisterCircle(i ICircle) {
	localCircle = i
}

var localCircle ICircle
