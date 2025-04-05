package like

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/like/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) LikeToggle(ctx context.Context, req *v1.LikeToggleReq) (res *v1.LikeToggleRes, err error) {
	// 判断用户是否登录
	if !service.Auth().LoginState(ctx) {
		return nil, gerror.New("用户未登录")
	}

	input := model.LikeToggleReq{
		TargetId:   req.TargetId,
		TargetType: req.TargetType,
	}

	output, err := service.Like().Toggle(ctx, input)
	if err != nil {
		return nil, err
	}

	return &v1.LikeToggleRes{
		LikeId:   output.LikeId,
		IsLiked:  output.IsLiked,
		LikeTime: output.LikeTime.String(),
	}, nil
}
