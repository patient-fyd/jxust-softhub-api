package auth

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"github.com/patient-fyd/jxust-softhub-api/internal/consts"
)

// GetLoginUserId 从上下文中获取当前登录用户ID
func GetLoginUserId(ctx context.Context) uint {
	// 从请求上下文中获取
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return 0
	}

	// 从上下文变量中获取用户ID
	userId := r.GetCtxVar(consts.CtxUserIdKey).Uint()
	return userId
}

// GetLoginUserName 从上下文中获取当前登录用户名
func GetLoginUserName(ctx context.Context) string {
	// 从请求上下文中获取
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return ""
	}

	// 从上下文变量中获取用户名
	userName := r.GetCtxVar(consts.CtxUserNameKey).String()
	return userName
}

// GetLoginUserRole 从上下文中获取当前登录用户角色ID
func GetLoginUserRole(ctx context.Context) uint {
	// 从请求上下文中获取
	r := g.RequestFromCtx(ctx)
	if r == nil {
		return 0
	}

	// 从上下文变量中获取用户角色
	roleId := r.GetCtxVar(consts.CtxUserRoleKey).Uint()
	return roleId
}

// IsAdmin 判断当前用户是否是管理员（角色ID为1、2、3）
func IsAdmin(ctx context.Context) bool {
	roleId := GetLoginUserRole(ctx)
	return roleId >= 1 && roleId <= 3
}

// CheckLogin 检查用户是否已登录
func CheckLogin(ctx context.Context) bool {
	return GetLoginUserId(ctx) > 0
}
