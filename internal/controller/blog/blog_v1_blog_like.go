package blog

import (
	"context"

	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogLike(ctx context.Context, req *v1.BlogLikeReq) (res *v1.BlogLikeRes, err error) {
	// 构建点赞参数
	input := model.BlogLikeInput{
		BlogId: req.BlogId,
	}

	// 调用Service层点赞博客
	err = service.Blog().Like(ctx, input)
	if err != nil {
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogLikeRes{}

	return res, nil
}
