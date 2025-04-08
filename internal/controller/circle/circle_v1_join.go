package circle

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/circle/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
	"github.com/patient-fyd/jxust-softhub-api/utility/auth"
)

// Join 关注或取消关注圈子
func (c *ControllerV1) Join(ctx context.Context, req *v1.JoinReq) (res *v1.JoinRes, err error) {
	// 判断用户是否已登录
	userId := auth.GetLoginUserId(ctx)
	if userId <= 0 {
		return nil, gerror.New("请先登录")
	}

	// 转换请求参数
	input := model.CircleJoinInput{
		CircleId: req.CircleId,
	}

	// 调用服务层
	output, err := service.Circle().Join(ctx, input)
	if err != nil {
		return nil, gerror.Wrap(err, "操作圈子关注状态失败")
	}

	// 转换响应参数
	res = &v1.JoinRes{
		IsFollowed: output.IsFollowed,
	}

	return res, nil
}
