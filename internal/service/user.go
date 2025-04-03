package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// IUser 用户服务接口
type IUser interface {
	// GetUserInfo 获取用户详细信息
	GetUserInfo(ctx context.Context, userId uint) (*model.UserDetailInfo, error)

	// GetUserList 获取用户列表
	GetUserList(ctx context.Context, in model.UserListInput) (*model.UserListOutput, error)

	// AssignRole 分配用户角色
	AssignRole(ctx context.Context, in model.UserAssignRoleInput) (*model.UserAssignRoleOutput, error)
}

var (
	localUser IUser
)

// User 获取用户服务实例
func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

// RegisterUser 注册用户服务实现
func RegisterUser(i IUser) {
	localUser = i
}
