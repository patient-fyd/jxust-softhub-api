package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// IMember 成员管理服务接口
type IMember interface {
	// List 获取成员列表
	List(ctx context.Context, in model.MemberListInput) (*model.MemberListOutput, error)

	// Detail 获取成员详情
	Detail(ctx context.Context, in model.MemberDetailInput) (*model.MemberDetailOutput, error)

	// Create 创建成员
	Create(ctx context.Context, in model.MemberCreateInput) (*model.MemberCreateOutput, error)

	// Update 更新成员信息
	Update(ctx context.Context, in model.MemberUpdateInput) (*model.MemberUpdateOutput, error)

	// Delete 删除成员
	Delete(ctx context.Context, in model.MemberDeleteInput) (*model.MemberDeleteOutput, error)
}

var (
	localMember IMember
)

// Member 获取成员管理服务实例
func Member() IMember {
	if localMember == nil {
		panic("implement not found for interface IMember, forgot register?")
	}
	return localMember
}

// RegisterMember 注册成员管理服务实现
func RegisterMember(i IMember) {
	localMember = i
}
