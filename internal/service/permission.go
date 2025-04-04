package service

import (
	"context"
)

// IPermission 权限服务接口
type IPermission interface {
	// HasPermission 检查当前登录用户是否拥有指定权限
	HasPermission(ctx context.Context, permissionId uint) (bool, error)

	// HasPermissionByKey 检查当前登录用户是否拥有指定权限键名
	HasPermissionByKey(ctx context.Context, permissionKey string) (bool, error)
}

var (
	localPermission IPermission
)

// Permission 获取权限服务实例
func Permission() IPermission {
	if localPermission == nil {
		panic("implement not found for interface IPermission, forgot register?")
	}
	return localPermission
}

// RegisterPermission 注册权限服务实现
func RegisterPermission(i IPermission) {
	localPermission = i
}
