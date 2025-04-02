package auth

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/auth/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	// 1. 转换请求参数
	loginInput := model.LoginInput{
		UserName: req.UserName,
		Password: req.Password,
	}

	// 2. 调用服务层处理登录
	loginOutput, err := service.Auth().Login(ctx, loginInput)
	if err != nil {
		return nil, err
	}

	// 3. 构造响应结果
	return &v1.LoginRes{
		Token: loginOutput.Token,
		User: v1.LoginUser{
			UserId:   loginOutput.User.UserId,
			UserName: loginOutput.User.UserName,
			Name:     loginOutput.User.Name,
		},
	}, nil
}
