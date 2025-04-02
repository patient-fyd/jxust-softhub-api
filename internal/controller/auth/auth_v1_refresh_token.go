package auth

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/auth/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	// 1. 转换请求参数
	refreshInput := model.TokenRefreshInput{
		Token: req.Token,
	}

	// 2. 调用服务层处理Token刷新
	refreshOutput, err := service.Auth().RefreshToken(ctx, refreshInput)
	if err != nil {
		return nil, err
	}

	// 3. 构造响应结果
	return &v1.RefreshTokenRes{
		Token: refreshOutput.Token,
	}, nil
}
