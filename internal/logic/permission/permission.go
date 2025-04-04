package permission

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/patient-fyd/jxust-softhub-api/internal/dao"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// 确保 sPermission 实现了 service.IPermission 接口
var _ service.IPermission = (*sPermission)(nil)

// sPermission 权限服务实现
type sPermission struct{}

// New 创建一个权限服务实例
func New() service.IPermission {
	return &sPermission{}
}

// init 初始化，注册服务
func init() {
	service.RegisterPermission(New())
}

// HasPermission 检查当前登录用户是否拥有指定权限
func (s *sPermission) HasPermission(ctx context.Context, permissionId uint) (bool, error) {
	// 获取当前登录用户ID
	userId := service.Auth().GetLoginUserId(ctx)
	if userId == 0 {
		return false, nil // 未登录用户没有任何权限
	}

	// 获取用户角色
	user, err := dao.Users.Ctx(ctx).Where("user_id", userId).One()
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil // 用户不存在
	}

	// 使用 gconv 包转换字段值
	roleId := gconv.Uint(user["role_id"])
	if roleId == 0 {
		return false, nil // 用户没有分配角色
	}

	// 超级管理员（角色ID为1）拥有所有权限
	if roleId == 1 {
		return true, nil
	}

	// 查询角色是否拥有指定权限
	count, err := dao.RolePermissions.Ctx(ctx).
		Where("role_id", roleId).
		Where("permission_id", permissionId).
		Count()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// HasPermissionByKey 检查当前登录用户是否拥有指定权限键名
func (s *sPermission) HasPermissionByKey(ctx context.Context, permissionKey string) (bool, error) {
	// 获取当前登录用户ID
	userId := service.Auth().GetLoginUserId(ctx)
	if userId == 0 {
		return false, nil // 未登录用户没有任何权限
	}

	// 获取用户角色
	user, err := dao.Users.Ctx(ctx).Where("user_id", userId).One()
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, nil // 用户不存在
	}

	// 使用 gconv 包转换字段值
	roleId := gconv.Uint(user["role_id"])
	if roleId == 0 {
		return false, nil // 用户没有分配角色
	}

	// 超级管理员（角色ID为1）拥有所有权限
	if roleId == 1 {
		return true, nil
	}

	// 先获取权限ID
	permission, err := dao.Permissions.Ctx(ctx).Where("permission_key", permissionKey).One()
	if err != nil {
		return false, err
	}
	if permission == nil {
		return false, nil // 权限不存在
	}

	permissionId := gconv.Uint(permission["permission_id"])
	if permissionId == 0 {
		return false, nil // 无效的权限ID
	}

	// 查询角色是否拥有指定权限
	count, err := dao.RolePermissions.Ctx(ctx).
		Where("role_id", roleId).
		Where("permission_id", permissionId).
		Count()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
