package auth

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/auth/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	// 调用业务层注册服务
	output, err := service.Auth().Register(ctx, model.UserRegisterInput{
		UserName: req.UserName,
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.RegisterRes{
		User: &v1.UserInfo{
			UserId:   output.UserId,
			UserName: output.UserName,
			Name:     output.Name,
			RoleId:   output.RoleId,
		},
	}, nil
}
