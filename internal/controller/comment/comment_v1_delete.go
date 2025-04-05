package comment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/comment/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) CommentDelete(ctx context.Context, req *v1.CommentDeleteReq) (res *v1.CommentDeleteRes, err error) {
	// 检查用户是否登录
	if !service.Auth().LoginState(ctx) {
		return nil, gerror.New("用户未登录")
	}

	input := model.CommentDeleteInput{
		CommentId: req.CommentId,
	}

	output, err := service.Comment().Delete(ctx, input)
	if err != nil {
		return nil, err
	}

	return &v1.CommentDeleteRes{
		Success: output.Success,
	}, nil
}
