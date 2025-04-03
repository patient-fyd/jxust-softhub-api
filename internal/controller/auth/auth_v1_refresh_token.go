package auth

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/auth/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	// 调用业务层Token刷新服务
	output, err := service.Auth().RefreshToken(ctx, model.RefreshTokenInput{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}

	// 返回结果
	return &v1.RefreshTokenRes{
		Token: output.Token,
	}, nil
}
