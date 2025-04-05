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
