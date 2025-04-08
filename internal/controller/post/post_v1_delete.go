package post

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/post/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	// 转换请求参数
	input := model.PostDeleteInput{
		PostId: req.PostId,
	}

	// 调用业务逻辑层
	output, err := service.Post().Delete(ctx, input)
	if err != nil {
		return nil, gerror.Wrap(err, "删除帖子失败")
	}

	// 转换响应参数
	return &v1.DeleteRes{
		Success: output.Success,
	}, nil
}
