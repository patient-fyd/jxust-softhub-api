package blog

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/model"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogRecommend(ctx context.Context, req *v1.BlogRecommendReq) (res *v1.BlogRecommendRes, err error) {
	// 构建推荐参数
	input := model.BlogRecommendInput{
		BlogId:      req.BlogId,
		IsRecommend: req.IsRecommend,
	}

	// 调用Service层设置博客推荐状态
	err = service.Blog().SetRecommend(ctx, input)
	if err != nil {
		g.Log().Error(ctx, "设置博客推荐状态失败:", err)
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogRecommendRes{
		Success: true,
	}

	return res, nil
}
