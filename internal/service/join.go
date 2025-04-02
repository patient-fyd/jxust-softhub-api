package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// IJoin 入会申请服务接口
type IJoin interface {
	// Apply 提交入会申请
	Apply(ctx context.Context, in model.JoinApplyInput) (*model.JoinApplyOutput, error)

	// List 获取入会申请列表
	List(ctx context.Context, in model.JoinListInput) (*model.JoinListOutput, error)

	// Detail 获取入会申请详情
	Detail(ctx context.Context, in model.JoinDetailInput) (*model.JoinDetailOutput, error)

	// Review 审核入会申请
	Review(ctx context.Context, in model.JoinReviewInput) (*model.JoinReviewOutput, error)
}

var (
	localJoin IJoin
)

// Join 获取入会申请服务实例
func Join() IJoin {
	if localJoin == nil {
		panic("implement not found for interface IJoin, forgot register?")
	}
	return localJoin
}

// RegisterJoin 注册入会申请服务实现
func RegisterJoin(i IJoin) {
	localJoin = i
}
