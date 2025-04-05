package follow

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/follow/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) FollowToggle(ctx context.Context, req *v1.FollowToggleReq) (res *v1.FollowToggleRes, err error) {
	// 判断用户是否登录
	if !service.Auth().LoginState(ctx) {
		return nil, gerror.New("用户未登录")
	}

	input := model.FollowToggleReq{
		FollowedId: req.FollowedId,
		FollowType: req.FollowType,
	}

	output, err := service.Follow().Toggle(ctx, input)
	if err != nil {
		return nil, err
	}

	return &v1.FollowToggleRes{
		FollowId:   output.FollowId,
		IsFollowed: output.IsFollowed,
		FollowTime: output.FollowTime.String(),
	}, nil
}
