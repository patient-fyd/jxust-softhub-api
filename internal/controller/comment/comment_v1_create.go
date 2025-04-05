package comment

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/comment/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) CommentCreate(ctx context.Context, req *v1.CommentCreateReq) (res *v1.CommentCreateRes, err error) {
	// 检查用户是否登录
	if !service.Auth().LoginState(ctx) {
		return nil, gerror.New("用户未登录")
	}

	input := model.CommentCreateInput{
		ContentType: req.ContentType,
		ContentId:   req.ContentId,
		Content:     req.Content,
	}

	output, err := service.Comment().Create(ctx, input)
	if err != nil {
		return nil, err
	}

	return &v1.CommentCreateRes{
		CommentId: output.CommentId,
	}, nil
}
