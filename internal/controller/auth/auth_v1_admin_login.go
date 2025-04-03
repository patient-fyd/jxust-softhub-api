package auth

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/auth/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) AdminLogin(ctx context.Context, req *v1.AdminLoginReq) (res *v1.AdminLoginRes, err error) {
	// 调用业务层管理员登录服务
	output, err := service.Auth().AdminLogin(ctx, model.AdminLoginInput{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.AdminLoginRes{
		Token: output.Token,
		User: &v1.UserInfo{
			UserId:   output.UserId,
			UserName: output.UserName,
			Name:     output.Name,
			RoleId:   output.RoleId,
		},
	}, nil
}
