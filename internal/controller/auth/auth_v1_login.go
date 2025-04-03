package auth

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/auth/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 调用业务层登录服务
	output, err := service.Auth().Login(ctx, model.UserLoginInput{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.LoginRes{
		Token: output.Token,
		User: &v1.UserInfo{
			UserId:   output.UserId,
			UserName: output.UserName,
			Name:     output.Name,
			RoleId:   output.RoleId,
		},
	}, nil
}
