package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogDelete(ctx context.Context, req *v1.BlogDeleteReq) (res *v1.BlogDeleteRes, err error) {
	// 构建删除参数
	input := model.BlogDeleteInput{
		BlogId: req.BlogId,
	}

	// 调用Service层删除博客
	err = service.Blog().Delete(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogDeleteRes{}

	return res, nil
}
