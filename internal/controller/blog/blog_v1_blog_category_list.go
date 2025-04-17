package blog

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogCategoryList(ctx context.Context, req *v1.BlogCategoryListReq) (res *v1.BlogCategoryListRes, err error) {
	// 调用Service层获取博客分类列表
	output, err := service.Blog().GetCategoryList(ctx)
	if err != nil {
		g.Log().Error(ctx, "获取博客分类列表失败:", err)
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogCategoryListRes{
		List: output.List,
	}

	return res, nil
}
