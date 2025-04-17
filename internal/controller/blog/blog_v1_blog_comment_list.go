package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogCommentList(ctx context.Context, req *v1.BlogCommentListReq) (res *v1.BlogCommentListRes, err error) {
	// 构建查询参数
	input := model.BlogCommentListInput{
		BlogId: req.BlogId,
		Page:   req.Page,
		Size:   req.Size,
	}

	// 调用Service层获取博客评论列表
	output, err := service.Blog().GetCommentList(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogCommentListRes{
		List:  output.List,
		Page:  output.Page,
		Size:  output.Size,
		Total: output.Total,
	}

	return res, nil
}
