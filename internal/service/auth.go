package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// IAuth 认证服务接口
type IAuth interface {
	Register(ctx context.Context, in model.RegisterInput) (*model.RegisterOutput, error)
	Login(ctx context.Context, in model.LoginInput) (*model.LoginOutput, error)
	RefreshToken(ctx context.Context, in model.TokenRefreshInput) (*model.TokenRefreshOutput, error)
}

var (
	localAuth IAuth
)

// Auth 获取认证服务实例
func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

// RegisterAuth 注册认证服务实现
func RegisterAuth(i IAuth) {
	localAuth = i
}
