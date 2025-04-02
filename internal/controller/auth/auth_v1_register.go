package auth

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/auth/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	// 1. 转换请求参数
	registerInput := model.RegisterInput{
		UserName: req.UserName,
		Password: req.Password,
		Name:     req.Name,
		Email:    req.Email,
		Phone:    req.Phone,
	}

	// 2. 调用服务层处理注册
	registerOutput, err := service.Auth().Register(ctx, registerInput)
	if err != nil {
		return nil, err
	}

	// 3. 构造响应结果
	return &v1.RegisterRes{
		UserId:     registerOutput.UserId,
		UserName:   registerOutput.UserName,
		CreateTime: registerOutput.CreateTime,
	}, nil
}
