package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogDetail(ctx context.Context, req *v1.BlogDetailReq) (res *v1.BlogDetailRes, err error) {
	// 构建查询参数
	input := model.BlogDetailInput{
		BlogId: req.BlogId,
	}

	// 调用Service层获取博客详情
	output, err := service.Blog().GetDetail(ctx, input)
	if err != nil {
		return nil, err
	}

	// 如果博客不存在
	if output == nil {
		return nil, nil
	}

	// 构建响应数据
	res = &v1.BlogDetailRes{
		BlogDetailOutput: output,
	}

	return res, nil
}
