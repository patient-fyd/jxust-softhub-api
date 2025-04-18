package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// 活动服务接口定义
type IActivity interface {
	// 活动管理
	Create(ctx context.Context, in model.ActivityCreateInput) (*model.ActivityCreateOutput, error)
	Update(ctx context.Context, in model.ActivityUpdateInput) (*model.ActivityUpdateOutput, error)
	Delete(ctx context.Context, in model.ActivityDeleteInput) error
	List(ctx context.Context, in model.ActivityListInput) (*model.ActivityListOutput, error)
	Detail(ctx context.Context, in model.ActivityDetailInput) (*model.ActivityDetailOutput, error)

	// 活动报名相关
	Register(ctx context.Context, in model.ActivityRegisterInput) (*model.ActivityRegisterOutput, error)
	RegisterList(ctx context.Context, in model.RegistrationListInput) (*model.RegistrationListOutput, error)
	ReviewRegistration(ctx context.Context, in model.RegistrationReviewInput) (*model.RegistrationReviewOutput, error)
}

var localActivity IActivity

// 获取活动服务实例
func Activity() IActivity {
	if localActivity == nil {
		panic("implement not found for interface IActivity, forgot register?")
	}
	return localActivity
}

// 注册活动服务实现
func RegisterActivity(i IActivity) {
	localActivity = i
}
