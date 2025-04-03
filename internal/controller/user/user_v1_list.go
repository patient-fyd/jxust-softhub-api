package user

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/user/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

// List 获取用户列表（管理员接口）
func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	// 从上下文中获取当前登录用户ID
	userId := service.Auth().GetLoginUserId(ctx)
	if userId == 0 {
		return nil, gerror.New("用户未登录或登录已过期")
	}

	// 获取当前用户信息
	userInfo, err := service.User().GetUserInfo(ctx, userId)
	if err != nil {
		return nil, err
	}

	// 检查是否有管理员权限（角色ID为1或2）
	if userInfo.RoleId != 1 && userInfo.RoleId != 2 {
		return nil, gerror.New("没有管理员权限")
	}

	// 调用业务层获取用户列表
	output, err := service.User().GetUserList(ctx, model.UserListInput{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Keyword:  req.Keyword,
	})
	if err != nil {
		return nil, err
	}

	// 转换为API响应格式
	var list []v1.UserInfo
	for _, user := range output.List {
		list = append(list, v1.UserInfo{
			UserId:    user.UserId,
			UserName:  user.UserName,
			Name:      user.Name,
			Email:     user.Email,
			Phone:     user.Phone,
			Avatar:    user.Avatar,
			RoleId:    user.RoleId,
			RoleName:  user.RoleName,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}

	return &v1.ListRes{
		List:     list,
		Total:    output.Total,
		PageNum:  output.PageNum,
		PageSize: output.PageSize,
	}, nil
}
