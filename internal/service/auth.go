package service

import (
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/model/entity"
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

	// GetLoginUserId 获取当前登录用户ID
	GetLoginUserId(ctx context.Context) uint

	// GetLoginUser 获取当前登录用户
	GetLoginUser(ctx context.Context) (*entity.Users, error)

	// LoginState 获取登录状态
	LoginState(ctx context.Context) bool

	// GenerateTokenFromDevice 生成设备令牌
	GenerateTokenFromDevice(ctx context.Context, userId int) (string, error)

	// GenerateToken 生成令牌
	GenerateToken(ctx context.Context, userId int, expire ...int) (string, error)

	// VerifyToken 验证令牌
	VerifyToken(ctx context.Context, token string) (*model.ClaimsOutput, error)

	// VerifyTokenMiddleware 令牌验证中间件
	VerifyTokenMiddleware(r *ghttp.Request)

	// CheckCapcha 检查验证码
	CheckCapcha(ctx context.Context, in model.CodeInput) bool

	// CheckPasswordInFile 检查密码是否在文件中
	CheckPasswordInFile(ctx context.Context, password string) bool

	// IsAdmin 检查是否为管理员
	IsAdmin(ctx context.Context) (bool, error)
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
