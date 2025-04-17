package blog

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/patient-fyd/jxust-softhub-api/api/blog/v1"
	"github.com/patient-fyd/jxust-softhub-api/internal/service"
)

func (c *ControllerV1) BlogTagList(ctx context.Context, req *v1.BlogTagListReq) (res *v1.BlogTagListRes, err error) {
	// 调用Service层获取博客标签列表
	output, err := service.Blog().GetTagList(ctx)
	if err != nil {
		g.Log().Error(ctx, "获取博客标签列表失败:", err)
		return nil, err
	}

	// 构建响应数据
	res = &v1.BlogTagListRes{
		List: output.List,
	}

	return res, nil
}
