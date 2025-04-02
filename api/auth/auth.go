package auth

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/auth/v1"
)

// IAuthV1 认证服务接口
type IAuthV1 interface {
	// Register 用户注册
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	// Login 用户登录
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	// RefreshToken 刷新Token
	RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error)
}
