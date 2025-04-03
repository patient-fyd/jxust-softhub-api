package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/user/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// Info 获取当前用户信息
func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	// 从上下文中获取当前登录用户ID
	userId := service.Auth().GetLoginUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("用户未登录或登录已过期")
	}

	// 获取用户信息
	userInfo, err := service.User().GetUserInfo(ctx, userId)
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.InfoRes{
		UserId:    userInfo.UserId,
		UserName:  userInfo.UserName,
		Name:      userInfo.Name,
		Email:     userInfo.Email,
		Phone:     userInfo.Phone,
		Avatar:    userInfo.Avatar,
		RoleId:    userInfo.RoleId,
		RoleName:  userInfo.RoleName,
		CreatedAt: userInfo.CreatedAt,
		UpdatedAt: userInfo.UpdatedAt,
	}, nil
}
