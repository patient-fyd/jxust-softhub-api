package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogUnlike(ctx context.Context, req *v1.BlogUnlikeReq) (res *v1.BlogUnlikeRes, err error) {
	// 构建取消点赞参数
	input := model.BlogUnlikeInput{
		BlogId: req.BlogId,
	}

	// 调用Service层取消点赞博客
	err = service.Blog().Unlike(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogUnlikeRes{}

	return res, nil
}
