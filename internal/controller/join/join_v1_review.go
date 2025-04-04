package join

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/join/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Review(ctx context.Context, req *v1.ReviewReq) (res *v1.ReviewRes, err error) {
	// 构造业务参数
	input := model.JoinReviewInput{
		ApplicationId: req.ApplicationId,
		Status:        req.Status,
		ReviewComment: req.ReviewComment,
	}

	// 调用服务
	_, err = service.Join().Review(ctx, input)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, err.Error())
	}

	// 返回结果
	return &v1.ReviewRes{}, nil
}
