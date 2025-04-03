package service

import (
	"context"

	"github.com/patient-fyd/jxust-softhub-api/internal/model"
)

// IAuth 认证服务接口
type IAuth interface {
	// Register 用户注册
	Register(ctx context.Context, in model.UserRegisterInput) (*model.UserRegisterOutput, error)

	// Login 用户登录
	Login(ctx context.Context, in model.UserLoginInput) (*model.UserLoginOutput, error)

	// AdminLogin 管理员登录
	AdminLogin(ctx context.Context, in model.AdminLoginInput) (*model.AdminLoginOutput, error)

	// RefreshToken 刷新Token
	RefreshToken(ctx context.Context, in model.RefreshTokenInput) (*model.RefreshTokenOutput, error)
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
